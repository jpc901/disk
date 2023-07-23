package service

import (
	"disk-server/master/service/system"
	"disk-server/master/service/user"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	UserServiceGroup   user.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
