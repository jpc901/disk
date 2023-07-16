package user

import (
	v1 "disk-master/api/v1"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouterV1 := Router.Group("v1/user")
	userApiV1 := v1.ApiGroupApp.UserApiGroup.UserApi
	{
		// v1
		userRouterV1.POST("signup", userApiV1.SignUp)
		userRouterV1.POST("signin", userApiV1.SignIn)
		userRouterV1.GET("info", userApiV1.GetUserInfo)
	}
}