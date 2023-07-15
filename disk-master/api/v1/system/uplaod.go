package system

import (
	"disk-master/model/enum"
	Err "disk-master/model/errors"
	"disk-master/model/request"
	"disk-master/model/response"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type UploadApi struct{}


func (up *UploadApi) LoadStatic(c *gin.Context) {
	response.BuildHtmlResponse(enum.UploadHtml, c)
}

func (up *UploadApi) UploadFile(c *gin.Context) {
	var requestData request.UploadFileRequest
	if err := c.ShouldBind(&requestData); err != nil {
		log.Error(err)
		response.BuildErrorResponse(err, c)
		return
	}

	if err := uploadService.UploadFile(requestData.File); err != nil {
		log.Error(err)
		response.BuildErrorResponse(Err.NewUploadFailedError("upload file failed"), c)
		return
	} else {
		response.BuildOkResponse(http.StatusOK, "upload file success", c)
	}
}
