package enrai

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sdslabs/gasper/configs"
	"github.com/sdslabs/gasper/types"
)

const (
	// DefaultServiceName is the name of the service proxying HTTP connections
	DefaultServiceName = types.Enrai

	// SSLServiceName is the name of the service proxying HTTPS connections
	SSLServiceName = types.EnraiSSL
)

// storage stores the reverse proxy records in the form of Key : Value pairs
// with Application Name as the key and its URL(IP:Port) as the value
var storage = types.NewRecordStorage()

// reverseProxy sets up the reverse proxy from the given domain to the target IP
func reverseProxy(c *gin.Context) {
	rootDomain := fmt.Sprintf(".%s", configs.GasperConfig.Domain)
	rootDomainWithPort := fmt.Sprintf("%s:%d", rootDomain, configs.ServiceConfig.Enrai.Port)

	if strings.HasSuffix(c.Request.Host, rootDomain) || strings.HasSuffix(c.Request.Host, rootDomainWithPort) {
		target, success := storage.Get(strings.Split(c.Request.Host, ".")[0])
		if !success {
			c.AbortWithStatusJSON(503, gin.H{
				"success": false,
				"message": "No such application exists",
			})
			return
		}

		director := func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = target
			req.Host = target
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
		return
	}

	c.AbortWithStatusJSON(403, gin.H{
		"success": false,
		"message": "Incorrect root domain",
	})
}

// NewService returns a new instance of the current microservice
func NewService() http.Handler {
	// router is the main routes handler for the current microservice package
	router := gin.Default()
	router.NoRoute(reverseProxy)
	return router
}
