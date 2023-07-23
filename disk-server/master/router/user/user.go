package user

import (
	v1 "master/api/v1"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouterV1 := Router.Group("v1/user")
	// userRouterV1.Use(middleware.JwtAuth())
	userApiV1 := v1.ApiGroupApp.UserApiGroup.UserApi
	{
		// v1
		userRouterV1.POST("signup", userApiV1.SignUp)
		userRouterV1.POST("signin", userApiV1.SignIn)
	}
}