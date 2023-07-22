package system

import (
	myDB "disk-master/dao/mysql"
	"disk-master/model"
	Err "disk-master/model/errors"
	"disk-master/model/request"

	log "github.com/sirupsen/logrus"
)

type FileMetaService struct {}
var FileMetaServiceApp = new(FileMetaService)


func (fm *FileMetaService) UpdateFile(updateRequest *request.FileUpdateRequest) (*model.FileMeta, error) {
	if updateRequest.OperateType != "0" {
		log.Error("operate type required is 0")
		return nil, Err.NewFileUpdateError("operate type required is 0")
	}
	// TODO: 要改成从缓存中取
	curFileMeta, err := fm.GetFileMeta(updateRequest.FileHash)
	if err != nil || curFileMeta == nil{
		log.Error("get file meta failed")
		return nil, err
	}
	curFileMeta.FileName = updateRequest.FileName
	// TODO: 要改成更新缓存从而更新用户表
	if err := fm.UpdateFileMetaDB(*curFileMeta); err != nil {
		log.Error(err)
		return nil, err
	}
	return curFileMeta, nil
}

// UpdateFileMetaDB:新增/更新文件元信息到mysql中
func (fm *FileMetaService) UpdateFileMetaDB(fileMeta model.FileMeta) error {
	return myDB.UpdateFile(fileMeta.FileSha1, fileMeta.FileName, fileMeta.FileSize, fileMeta.Location)
}

// 更新文件原信息到mysql中
func (fm *FileMetaService) UpdateUserFileMetaDB(username, fileSha1, fileName string, fileSize int64) error {
	return myDB.UpdateUserFile(username, fileSha1, fileName, fileSize)
}


func (fm *FileMetaService) GetFileMeta(fileSha1 string) (*model.FileMeta, error) {
	tmpFile, err := myDB.GetFileMeta(fileSha1)
	if tmpFile == nil || err != nil {
		log.Info("没有任何记录")
		return nil, err
	}
	fileMeta := &model.FileMeta{
		FileSha1: tmpFile.FileHash,
		FileName: tmpFile.FileName.String,
		FileSize: tmpFile.FileSize.Int64,
		Location: tmpFile.FileAddr.String,
	}
	return fileMeta, nil
}

func (fm *FileMetaService) GetUserFileMeta(uid int64) ([]*model.UserFileMeta, error) {
	username, err := myDB.GetUsername(uid)
	if err != nil {
		return nil, err
	}
	return myDB.QueryUserFileMeta(username)
}

func (fm *FileMetaService) DeleteUserFile(fileHash, fileName string) (error) {
	err := myDB.DeleteUserFile(fileHash, fileName)
	if err != nil {
		return err
	}
	return nil
}

