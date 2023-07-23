package system

import (
	v1 "master/api/v1"
	"master/middleware"

	"github.com/gin-gonic/gin"
)

type UploadRouter struct{}

func (up *UploadRouter) InitUploadRouter(Router *gin.RouterGroup) {
	uploadRouterV1 := Router.Group("v1/file")
	uploadRouterV1.Use(middleware.JwtAuth())
	uploadApiV1 := v1.ApiGroupApp.UploadApiGroup.UploadApi
	{
		// v1
		uploadRouterV1.POST("upload", uploadApiV1.UploadFile)
		uploadRouterV1.POST("fastupload", uploadApiV1.FastUploadFile)
		uploadRouterV1.POST("mpupload/init", uploadApiV1.MpUploadFileInit)
		uploadRouterV1.POST("mpupload/chunk", uploadApiV1.UploadPart)
		uploadRouterV1.GET("mpupload/check", uploadApiV1.CheckChunkExist)
		uploadRouterV1.POST("mpupload/merge", uploadApiV1.MergeChunk)
	}
}