package service

import (
	"master/service/system"
	"master/service/user"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	UserServiceGroup   user.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
