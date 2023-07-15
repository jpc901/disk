package system

import (
	Err "disk-master/model/errors"
	"disk-master/model/response"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type UploadApi struct{}

const (
	uploadHtml = "upload.html"
	uploadPath = "./static/tmpfile"
)

func (up *UploadApi) LoadStatic(c *gin.Context) {
	c.HTML(http.StatusOK, uploadHtml, nil)
}

func (up *UploadApi) UploadFile(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		log.Error(err)
		response.BuildErrorResponse(Err.NewUploadFailedError("upload failed of param"), c)
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		log.Error(err)
		response.BuildErrorResponse(Err.NewUploadFailedError("upload failed of open file"), c)
		return
	}

	defer file.Close()

	filePath := uploadPath + fileHeader.Filename
	log.Info(filePath)
	newFile, err := os.Create(filePath)
	if err != nil {
		log.Error(err)
		response.BuildErrorResponse(Err.NewUploadFailedError("upload failed of create file"), c)
		return
	}
	defer newFile.Close()

	io.Copy(newFile, file)

	response.BuildOkResponse(http.StatusOK, "success", c)
}