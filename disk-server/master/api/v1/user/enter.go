package user

import "disk-server/master/service"

type ApiGroup struct {
	UserApi
}

var (
	userService = service.ServiceGroupApp.UserServiceGroup.UserService
)