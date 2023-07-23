package system

import (
	v1 "master/api/v1"
	"master/middleware"

	"github.com/gin-gonic/gin"
)

type FileOperateRouter struct {}

func (fo *FileOperateRouter)InitFileOperateRouter(Router *gin.RouterGroup) {
	fileOperateRouterV1 := Router.Group("v1/file")
	fileOperateRouterV1.Use(middleware.JwtAuth())
	fileOperateApiV1 := v1.ApiGroupApp.UploadApiGroup.FileOperateApi
	{
		fileOperateRouterV1.GET("meta", fileOperateApiV1.GetFileMeta)
		fileOperateRouterV1.POST("download", fileOperateApiV1.FileDownload)
		fileOperateRouterV1.DELETE("delete", fileOperateApiV1.FileDelete)
		fileOperateRouterV1.PUT("update", fileOperateApiV1.FileUpdate)
		fileOperateRouterV1.GET("query", fileOperateApiV1.QueryUserFile)
	}
}