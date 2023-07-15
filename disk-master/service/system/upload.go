package system

import (
	"disk-master/model/enum"
	"io"
	"mime/multipart"
	"os"

	log "github.com/sirupsen/logrus"
)

type UploadService struct{}



func (up *UploadService) UploadFile(fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		log.Errorf("file open failed, error: %v",err)
		return err
	}

	defer file.Close()

	filePath := enum.UploadPath + fileHeader.Filename
	newFile, err := os.Create(filePath)
	if err != nil {
		log.Errorf("upload failed of create file, error: %v", err)
		return err
	}
	defer newFile.Close()
	io.Copy(newFile, file)
	return nil
}