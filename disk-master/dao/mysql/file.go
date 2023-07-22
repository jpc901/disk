package db

import (
	"database/sql"
	"disk-master/global"

	log "github.com/sirupsen/logrus"
)

// TableFile : 文件表结构体
type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

// OnFileUploadFinished : 文件上传完成,保存meta
func UpdateFile(fileHash string, fileName string, fileSize int64, fileAddr string) error {

	stmt, err := global.DB.GetConn().Prepare("insert ignore into tbl_file" +
		"(`file_sha1`,`file_name`,`file_size`,`file_addr`,`status`) values(?,?,?,?,1)")
	if err != nil {
		log.Errorf("Failed to prepare statement,err: %s",  err.Error())
		return err
	}
	defer stmt.Close()
	ret, err := stmt.Exec(fileHash, fileName, fileSize, fileAddr)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	if rf, err := ret.RowsAffected(); err == nil {
		if rf <= 0 {
			log.Info("File with hash been upload before", fileHash)
		}
	}
	return err
}

// GetFileMeta : 从mysql获取文件元信息
func GetFileMeta(fileHash string) (*TableFile, error) {
	stmt, err := global.DB.GetConn().Prepare("select file_sha1,file_addr,file_name,file_size from tbl_file where file_sha1=? and status=1 limit 1")
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer stmt.Close()

	tmpFile := TableFile{}
	err = stmt.QueryRow(fileHash).Scan(&tmpFile.FileHash, &tmpFile.FileAddr, &tmpFile.FileName, &tmpFile.FileSize)
	if err != nil {
		if err == sql.ErrNoRows {
			// 查不到对应记录， 返回参数及错误均为nil
			return nil, nil
		}
		log.Error(err.Error())
		return nil, err
	}
	return &tmpFile, nil
}


func DeleteUserFile(fileHash, fileName string) error {
	stmt, err := global.DB.GetConn().Prepare("delete from tbl_user_file where file_sha1=? and file_name=? limit 1")
	if err != nil {
		log.Error(err.Error())
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(fileHash, fileName)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}