package service

import (
	"disk-master/service/system"
	"disk-master/service/user"
)

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
	UserServiceGroup   user.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
