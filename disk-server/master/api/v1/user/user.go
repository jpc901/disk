package user

import (
	"context"
	"net/http"

	"github.com/jpc901/disk-common/model"
	Err "github.com/jpc901/disk-common/model/errors"
	"github.com/jpc901/disk-common/model/request"
	"github.com/jpc901/disk-common/model/response"

	"disk-server/master/grpc/account"
	accountClientPb "disk-server/master/grpc/proto/account"

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
	req := &accountClientPb.UserSignUpRequest{
		Username: requestData.Username,
		Password: requestData.Password,
		ConfirmPassword: requestData.Password,
	}

	resp, err := account.AccountClient.SignUp(context.Background(), req)
	if resp.Code != 0 || err != nil {
		response.BuildErrorResponse(Err.NewUserSignUpError("注册失败"), c)
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
	req := &accountClientPb.UserSignInRequest{
		Username: requestData.Username,
		Password: requestData.Password,
	}
	resp, err := account.AccountClient.SignIn(context.Background(), req)
	if resp.Code != 0 || err != nil {
		response.BuildErrorResponse(Err.NewUserSignInError("登录失败"), c)
		return
	}
	user := &model.User{
		Uid: resp.Uid,
		Username: resp.Username,
		SignUpAt: resp.SignUpAt,
	}
	data := &response.UserLogin{
		Token: resp.Token,
		User: *user,
	}
	response.BuildOkResponse(http.StatusOK, data, c)
}