package db

import (
	"disk-master/global"
	"disk-master/model"
	"time"

	log "github.com/sirupsen/logrus"
)

// 更新用户文件表
func UpdateUserFile(username, fileHash, filename string, fileSize int64) error {
	sqlStr := "insert ignore into tbl_user_file (`user_name`,`file_sha1`,`file_name`,`file_size`,`upload_at`) values (?,?,?,?,?)"
	stmt, err := global.DB.GetConn().Prepare(sqlStr)
	if err != nil {
		log.Error(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, fileHash, filename, fileSize, time.Now())
	if err != nil {
		return err
	}
	return nil
}

func QueryUserFileMeta(username string) ([]*model.UserFileMeta, error) {
	sqlStr := "select file_sha1,file_name,file_size,upload_at, last_update from tbl_user_file where user_name=?"
	stmt, err := global.DB.GetConn().Prepare(sqlStr)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(username)
	if err != nil {
		return nil, err
	}

	var userFiles []*model.UserFileMeta
	for rows.Next() {
		userFile := &model.UserFileMeta{}
		err = rows.Scan(&userFile.FileHash, &userFile.FileName, &userFile.FileSize, &userFile.UploadAt, &userFile.LastUpdated)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		userFile.UserName = username
		userFiles = append(userFiles, userFile)
	}
	return userFiles, nil
}