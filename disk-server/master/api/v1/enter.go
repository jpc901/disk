package v1

import (
	"master/api/v1/health"
	"master/api/v1/system"
	"master/api/v1/user"
)

type ApiGroup struct{
	HealthApiGroup health.ApiGroup
	UploadApiGroup system.ApiGroup
	UserApiGroup   user.ApiGroup
}

var ApiGroupApp = new(ApiGroup)