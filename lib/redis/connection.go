package redis

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
	"github.com/sdslabs/gasper/configs"
	"github.com/sdslabs/gasper/lib/utils"
)

var client = redis.NewClient(&redis.Options{
	Addr:     fmt.Sprintf("%s:%d", configs.RedisConfig.Host, configs.RedisConfig.Port),
	Password: configs.RedisConfig.Password,
	DB:       configs.RedisConfig.DB,
})

func init() {
	_, err := client.Ping().Result()
	if err != nil {
		utils.Log("Redis connection was not established", utils.ErrorTAG)
		utils.LogError(err)
		os.Exit(1)
	} else {
		utils.LogInfo("Redis Connection Established")
	}
}
