package router

import (
	"disk-master/router/health"
	"disk-master/router/system"
	"disk-master/router/user"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Routers struct {
	Health health.RouterGroup
	System system.RouterGroup
	User   user.RouterGroup
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
	userRouter := RouterGroupApp.User
	// out api
	OutRouter := Router.Group("api")
	{
		systemRouter.InitUploadRouter(OutRouter)
		systemRouter.InitFileOperateRouter(OutRouter)
		userRouter.InitUserRouter(OutRouter)
	}

	// in api
	InRouter := Router.Group("")
	{
		healthRouter.InitHealthyRouter(InRouter)
	}
	log.Info("successfully initialized route group")
	return Router
}
