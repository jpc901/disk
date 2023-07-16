package user

import "disk-master/service"

type ApiGroup struct {
	UserApi
}

var (
	userService = service.ServiceGroupApp.UserServiceGroup.UserService
)