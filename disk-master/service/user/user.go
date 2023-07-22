package user

import (
	"time"

	myDB "github.com/jpc901/disk-common/mapper"
	"github.com/jpc901/disk-common/model"
	Err "github.com/jpc901/disk-common/model/errors"
	"github.com/jpc901/disk-common/model/request"
	"github.com/jpc901/disk-common/model/response"

	"github.com/jpc901/disk-common/jwt"
	"github.com/jpc901/disk-common/snowflake"
	util "github.com/jpc901/disk-common/utils"
	log "github.com/sirupsen/logrus"
)

type UserService struct {}


func (u *UserService) SignUp(user *request.UserSignUpRequest) error {

	if err := myDB.CheckUserExist(user.Username); err != nil {
		return err
	}
	// 生成 uuid
	uuid, err := snowflake.GetId(1, 1)
	if err != nil {
		log.Error("gen uuid failed, err:", err)
		return err
	}
	// 对密码进行加盐及取Sha1值加密
	encPasswd := util.Sha1([]byte(user.Password + model.PwdSalt))
	if err := myDB.UserSignUp(uuid, user.Username, encPasswd); err != nil {
		log.Error("sign up failed, err", err)
		return Err.NewUserSignUpError("注册失败")
	}
	return nil
}

func (u *UserService) SignIn(user *request.UserSignInRequest) (*response.UserLogin, error) {

	// 获取加密后的密码
	user.Password = util.Sha1([]byte(user.Password + model.PwdSalt))

	// 获取用户信息
	userInfo, err := myDB.GetUserInfo(user.Username)
	if err != nil {
		log.Warn("get user info  failed err:", err)
		return nil, err
	}
	if userInfo == nil {
		return nil, Err.NewUserNotExistError("用户不存在")
	}
	if userInfo.Password != user.Password {
		return nil, Err.NewPasswordError("密码错误")
	}

	// 生成token
	token, err := jwt.CreateToken(userInfo.Uid, time.Now().Add(24*60*60*time.Second).Unix())
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &response.UserLogin{
		Token: token,
		User: *userInfo,
	}, nil
}