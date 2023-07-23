package handler

import (
	"context"
	"time"

	account "disk-server/account/proto"

	"github.com/jpc901/disk-common/jwt"
	myDB "github.com/jpc901/disk-common/mapper"
	"github.com/jpc901/disk-common/model"
	Err "github.com/jpc901/disk-common/model/errors"
	"github.com/jpc901/disk-common/snowflake"
	util "github.com/jpc901/disk-common/utils"
	log "github.com/sirupsen/logrus"
)

type User struct{
	account.UnimplementedUserServiceServer
}

func (u *User) SignUp(ctx context.Context,  req *account.UserSignUpRequest) (*account.UserSignUpResponse, error) {
	if err := myDB.CheckUserExist(req.Username); err != nil {
		return nil, err
	}
	// 生成 uuid
	uuid, err := snowflake.GetId(1, 1)
	if err != nil {
		log.Error("gen uuid failed, err:", err)
		return nil, err
	}
	// 对密码进行加盐及取Sha1值加密
	encPasswd := util.Sha1([]byte(req.Password + model.PwdSalt))
	if err := myDB.UserSignUp(uuid, req.Username, encPasswd); err != nil {
		log.Error("sign up failed, err", err)
		return nil, Err.NewUserSignUpError("注册失败")
	}
	return &account.UserSignUpResponse{
		Code: 0,
	}, nil
}

func (u *User) SignIn(ctx context.Context, req *account.UserSignInRequest) (*account.UserSignInResponse, error) {
		// 获取加密后的密码
	req.Password = util.Sha1([]byte(req.Password + model.PwdSalt))

	// 获取用户信息
	userInfo, err := myDB.GetUserInfo(req.Username)
	if err != nil {
		log.Warn("get user info  failed err:", err)
		return nil, err
	}
	if userInfo == nil {
		return nil, Err.NewUserNotExistError("用户不存在")
	}
	if userInfo.Password != req.Password {
		return nil, Err.NewPasswordError("密码错误")
	}

	// 生成token
	token, err := jwt.CreateToken(userInfo.Uid, time.Now().Add(24*60*60*time.Second).Unix())
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &account.UserSignInResponse{
		Code: 0,
		Token: token,
		Uid: userInfo.Uid,
		SignUpAt: userInfo.SignUpAt,
		Username: userInfo.Username,
	}, nil
}