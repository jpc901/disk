package system

import (
	myDB "disk-master/dao/mysql"
	"disk-master/model"
)

type FileMetaService struct {}
var FileMetaServiceApp = new(FileMetaService)

var fileMetaMap map[string]model.FileMeta

func init() {
	fileMetaMap = make(map[string]model.FileMeta)
}

func (fm *FileMetaService) UpdateFileMeta(fileMeta model.FileMeta) {
	fileMetaMap[fileMeta.FileSha1] = fileMeta
}

// UpdateFileMetaDB:新增/更新文件元信息到mysql中
func (fm *FileMetaService) UpdateFileMetaDB(fileMeta model.FileMeta) error {
	return myDB.OnFileUploadFinished(fileMeta.FileSha1, fileMeta.FileName, fileMeta.FileSize, fileMeta.Location)
}

//GetFileMetaDB:从mysql获取文件元信息
func (fm *FileMetaService) GetFileMetaDB(fileSha1 string) (*model.FileMeta, error) {
	tmpFile, err := myDB.GetFileMeta(fileSha1)
	if tmpFile == nil || err != nil {
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

// GetFileMeta:通过Sha1获取文件的元信息对象
func (fm *FileMetaService) GetFileMeta(fileSha1 string) model.FileMeta {
	return fileMetaMap[fileSha1]
}

// RemoveFileMeta : 删除元信息
func (fm *FileMetaService) RemoveFileMeta(fileSha1 string) {
	delete(fileMetaMap, fileSha1)
}
