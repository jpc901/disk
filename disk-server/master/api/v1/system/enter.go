package system

import "disk-server/master/service"

type ApiGroup struct {
	UploadApi
	FileOperateApi
}

var (
	uploadService = service.ServiceGroupApp.SystemServiceGroup.UploadService
	fileMetaService = service.ServiceGroupApp.SystemServiceGroup.FileMetaService
)
