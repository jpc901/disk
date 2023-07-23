package user

import "master/service"

type ApiGroup struct {
	UserApi
}

var (
	userService = service.ServiceGroupApp.UserServiceGroup.UserService
)