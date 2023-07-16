package v1

import (
	"disk-master/api/v1/health"
	"disk-master/api/v1/system"
	"disk-master/api/v1/user"
)

type ApiGroup struct{
	HealthApiGroup health.ApiGroup
	UploadApiGroup system.ApiGroup
	UserApiGroup   user.ApiGroup
}

var ApiGroupApp = new(ApiGroup)