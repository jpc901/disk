package v1

import (
	"disk-master/api/v1/health"
	"disk-master/api/v1/system"
)

type ApiGroup struct{
	HealthApiGroup health.ApiGroup
	UploadApiGroup system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)