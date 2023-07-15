package system

import (
	"disk-master/model"
	"disk-master/model/enum"
	"io"
	"mime/multipart"
	"os"
	"time"

	util "github.com/jpc901/disk-common/utils"
	log "github.com/sirupsen/logrus"
)

type UploadService struct{}



func (up *UploadService) UploadFile(fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		log.Errorf("file open failed, error: %v",err)
		return err
	}
	fileMeta := model.FileMeta{
		FileName: fileHeader.Filename,
		Location: enum.UploadPath + fileHeader.Filename,
		UploadAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	defer file.Close()

	newFile, err := os.Create(fileMeta.Location)
	if err != nil {
		log.Errorf("Upload failed of create file, error: %v", err)
		return err
	}
	defer newFile.Close()

	fileMeta.FileSize, err = io.Copy(newFile, file)
	if err != nil {
		log.Errorf("Failed to save data into file,err: %v", err)
		return err
	}

	newFile.Seek(0, 0)
	fileMeta.FileSha1 = util.FileSha1(newFile)

	err = FileMetaServiceApp.UpdateFileMetaDB(fileMeta)
	if err == nil {
		log.Info("success [^_^]")
	}
	return err
}