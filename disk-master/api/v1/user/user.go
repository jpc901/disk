package user

import (
	"disk-master/model/request"
	"disk-master/model/response"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type UserApi struct{}

func (u *UserApi) SignUp(c *gin.Context) {
	var requestData request.UserSignUpRequest
	if err := c.ShouldBind(&requestData); err != nil {
		log.Error(err)
		response.BuildErrorResponse(err, c)
		return
	}
	if err := userService.SignUp(&requestData); err != nil {
		log.Error("sign up failed, err: ",err)
		response.BuildErrorResponse(err, c)
		return
	}
	response.BuildOkResponse(http.StatusOK, "sign up success", c)
}

func (u *UserApi) SignIn(c *gin.Context) {
	var requestData request.UserSignInRequest
	if err := c.ShouldBind(&requestData); err != nil {
		log.Error(err)
		response.BuildErrorResponse(err, c)
		return
	}
	data, err := userService.SignIn(&requestData)
	if err != nil {
		log.Error("sign in failed, err: ",err)
		response.BuildErrorResponse(err, c)
		return
	}
	response.BuildOkResponse(http.StatusOK, data, c)
}