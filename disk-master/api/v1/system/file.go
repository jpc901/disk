package system

import (
	Err "disk-master/model/errors"
	"disk-master/model/request"
	"disk-master/model/response"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type FileOperateApi struct{}

func (fo *FileOperateApi) GetFileMeta(c *gin.Context) {
	var requestData request.FileMetaRequest
	if err := c.ShouldBind(&requestData); err != nil {
		log.Error(err)
		response.BuildErrorResponse(err, c)
		return
	}
	data, err := fileMetaService.GetFileMeta(requestData.FileHash)
	if err != nil {
		log.Error(err)
		response.BuildErrorResponse(Err.NewGetFileMetaError("get file meta from db failed"), c)
		return
	}
	response.BuildOkResponse(http.StatusOK, data, c)
}

func (fo *FileOperateApi) FileDownload(c *gin.Context) {
	var requestData request.FileDownloadRequest
	if err := c.ShouldBind(&requestData); err != nil {
		log.Error(err)
		response.BuildErrorResponse(err, c)
		return
	}
	// 实际改成从环从中取，当前用户表
	// GetFileMeta
	fileMeta, err := fileMetaService.GetFileMeta(requestData.FileHash)
	if err != nil {
		log.Error("get file meta failed,  err: ",err)
		response.BuildErrorResponse(err, c)
		return
	}
	file, err := os.Open(fileMeta.Location)
	if err != nil {
		log.Error("open file failed,  err: ",err)
		response.BuildErrorResponse(err, c)
		return
	}
	defer file.Close()
	header := c.Writer.Header()
	header["Content-type"] = []string{"application/octet-stream"}
	header["Content-Disposition"] = []string{"attachment; filename= " + fileMeta.FileName}
	io.Copy(c.Writer, file)
	response.BuildOkResponse(http.StatusOK, nil, c)
}

// 未改
func (fo *FileOperateApi) FileDelete(c *gin.Context) {
	var requestData request.FileDeleteRequest
	if err := c.ShouldBindJSON(&requestData); err != nil {
		log.Error(err)
		response.BuildErrorResponse(err, c)
		return
	}
	err := fileMetaService.DeleteUserFile(requestData.FileHash, requestData.FileName)
	if err != nil {
		log.Error(err)
		response.BuildErrorResponse(err, c)
		return
	}
	response.BuildOkResponse(http.StatusOK, nil, c)
}

// 未改
func (fo *FileOperateApi) FileUpdate(c *gin.Context) {
	var requestData request.FileUpdateRequest
	if err := c.ShouldBind(&requestData); err != nil {
		log.Error(err)
		response.BuildErrorResponse(err, c)
		return
	}

	data, err := fileMetaService.UpdateFile(&requestData)
	if err != nil {
		log.Error(err)
		response.BuildErrorResponse(Err.NewFileUpdateError("update file failed"), c)
		return
	}
	response.BuildOkResponse(http.StatusOK, data, c)
}

func (fo *FileOperateApi) QueryUserFile(c *gin.Context) {
	uidAny, _ := c.Get("uid")
	uid := uidAny.(int64)
	data, err := fileMetaService.GetUserFileMeta(uid)
	if err != nil {
		log.Error(err)
		response.BuildErrorResponse(Err.NewGetUserFileError("获取用户文件失败"), c)
		return
	}
	response.BuildOkResponse(http.StatusOK, data, c)
}