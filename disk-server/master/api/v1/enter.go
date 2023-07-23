package v1

import (
	"disk-server/master/api/v1/health"
	"disk-server/master/api/v1/system"
	"disk-server/master/api/v1/user"
)

type ApiGroup struct{
	HealthApiGroup health.ApiGroup
	UploadApiGroup system.ApiGroup
	UserApiGroup   user.ApiGroup
}

var ApiGroupApp = new(ApiGroup)