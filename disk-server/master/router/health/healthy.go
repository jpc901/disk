package health

import (
	v1 "master/api/v1"

	"github.com/gin-gonic/gin"
)

type HealthyRouter struct{}

func (h *HealthyRouter) InitHealthyRouter(Router *gin.RouterGroup) {
	healthyRouter := Router.Group("healthy")
	healthyApi := v1.ApiGroupApp.HealthApiGroup.HealthyApi
	{
		healthyRouter.GET("readiness", healthyApi.Readiness)
		healthyRouter.GET("liveness", healthyApi.Liveness)
	}

}
