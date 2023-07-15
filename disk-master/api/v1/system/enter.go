package system

import "disk-master/service"

type ApiGroup struct{
	UploadApi
}

var (
	uploadService         = service.ServiceGroupApp.SystemServiceGroup.UploadService
)