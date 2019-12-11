package main

import (
	"os"
	"strings"

	"github.com/sdslabs/gasper/configs"
	"github.com/sdslabs/gasper/lib/utils"
	"github.com/sdslabs/gasper/services/enrai"
	"github.com/sdslabs/gasper/services/hikari"
	"github.com/sdslabs/gasper/services/kaze"
	"golang.org/x/sync/errgroup"
)

func initKaze() {
	go kaze.ScheduleServiceExposure()
	if configs.ServiceConfig.Kaze.Deploy {
		go kaze.ScheduleCleanup()
		go kaze.ScheduleCollectMetrics()
	}
}

func initHikari() {
	if configs.ServiceConfig.Hikari.Deploy {
		go hikari.ScheduleUpdate()
	}
}

func initEnrai() {
	if configs.ServiceConfig.Enrai.Deploy {
		go enrai.ScheduleUpdate()
	}
}

func initServices() {
	var g errgroup.Group
	for service, launcher := range launcherBindings {
		if launcher.Deploy {
			g.Go(launcher.Start)
			utils.LogInfo("%s Service Active\n", strings.Title(service))
		}
	}
	if err := g.Wait(); err != nil {
		utils.LogError(err)
		os.Exit(1)
	}
}

func main() {
	initKaze()
	initHikari()
	initEnrai()
	initServices()
}
