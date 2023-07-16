package db

import (
	"disk-master/global"
	"disk-master/model"

	Err "disk-master/model/errors"

	log "github.com/sirupsen/logrus"
)

// UserSignUp : 通过用户名及密码完成user表的注册操作
func UserSignUp(uuid int64, username string, passwd string) error {
	sqlStr := "insert ignore into tbl_user (`uid`,`user_name`,`user_pwd`) values (?,?,?)"
	stmt, err := global.DB.GetConn().Prepare(sqlStr)
	if err != nil {
		log.Error("Failed to insert, err:" + err.Error())
		return err
	}
	defer stmt.Close()
	ret, err := stmt.Exec(uuid, username, passwd)
	if err != nil {
		log.Error("Failed to insert, err:" + err.Error())
		return err
	}
	if rowsAffected, err := ret.RowsAffected(); nil == err && rowsAffected > 0 {
		return nil
	}
	return err
}


// UserSignin : 判断密码是否一致
func UserSignIn(username string, encpwd string) error {
	stmt, err := global.DB.GetConn().Prepare("select * from tbl_user where user_name=? limit 1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(username)
	if err != nil {
		return err
	} else if rows == nil {
		return err
	}
	return err
}

// UpdateToken : 刷新用户登录的token
func UpdateToken(username string, token string) error {
	stmt, err := global.DB.GetConn().Prepare(
		"replace into tbl_user_token (`user_name`,`user_token`) values (?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, token)
	if err != nil {
		return err
	}
	return nil
}

// GetUserInfo : 查询用户信息
func GetUserInfo(username string) (userInfo *model.User, err error) {
	userInfo = &model.User{}
	sqlStr := "select uid, user_name, signup_at, user_pwd from tbl_user where user_name=? limit 1"
	stmt, err := global.DB.GetConn().Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&userInfo.Uid, &userInfo.Username, &userInfo.SignUpAt, &userInfo.Password)
	if err != nil {
		return nil, err
	}
	return
}

// CheckUserExist 检查制定用户名的用户是否存在
func CheckUserExist(username string) (error) {
	sqlStr := `select count(uid) from tbl_user where user_name = ?`
	stmt, err := global.DB.GetConn().Prepare(sqlStr)
	if err != nil {
		log.Error(err)
		return err
	}
	defer stmt.Close()
	var count int
	if err = stmt.QueryRow(username).Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		return Err.NewUserExistError("用户已存在")
	}
	return err
}