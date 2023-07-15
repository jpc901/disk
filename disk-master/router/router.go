package router

import (
	"disk-master/router/health"
	"disk-master/router/system"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)


type Routers struct {
	Health health.RouterGroup
	System system.RouterGroup
}

var RouterGroupApp = new(Routers)

const templatesMatchPath = "static/view/*"

func Init() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{
			"/healthy/readiness",
			"/healthy/liveness",
		},
	}), gin.Recovery())
	pprof.Register(Router)

	// load static
	Router.LoadHTMLGlob(templatesMatchPath)
	Router.Static("/static", "./static")

	systemRouter := RouterGroupApp.System
	healthRouter := RouterGroupApp.Health

	// out api
	OutRouter := Router.Group("api")
	{
		systemRouter.InitUploadRouter(OutRouter)
	}

	// in api
	InRouter := Router.Group("")
	{
		healthRouter.InitHealthyRouter(InRouter)
	}
	log.Info("successfully initialized route group")
	return Router
}