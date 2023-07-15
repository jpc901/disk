package system

import (
	v1 "disk-master/api/v1"

	"github.com/gin-gonic/gin"
)

type UploadRouter struct{}

func (up *UploadRouter) InitUploadRouter(Router *gin.RouterGroup) {
	uploadRouterV1 := Router.Group("v1/file")
	uploadApiV1 := v1.ApiGroupApp.UploadApiGroup.UploadApi
	{
		// v1
		uploadRouterV1.GET("upload", uploadApiV1.LoadStatic)
		uploadRouterV1.POST("upload", uploadApiV1.UploadFile)
	}

}