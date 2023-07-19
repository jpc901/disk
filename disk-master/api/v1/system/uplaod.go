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

	if err := uploadService.UploadFile(requestData.Username, requestData.File); err != nil {
		log.Error(err)
		response.BuildErrorResponse(Err.NewUploadFailedError("upload file failed"), c)
		return
	} else {
		response.BuildOkResponse(http.StatusOK, "upload file success", c)
	}
}

func (up *UploadApi) FastUploadFile(c *gin.Context) {
	var requestData request.FastUploadRequest
	if err := c.ShouldBind(&requestData); err != nil {
		log.Error(err)
		response.BuildErrorResponse(err, c)
		return
	}
	// 2. 从文件表中查询相同hash的文件记录
	fileMate, err := fileMetaService.GetFileMetaDB(requestData.FileHash)
	if err != nil {
		log.Error(err)
		response.BuildErrorResponse(Err.NewFastUploadError("查询相同文件信息出错"), c)
		return
	}
	if fileMate == nil {
		response.BuildErrorResponse(Err.NewFastUploadError("没有相同文件信息，请采用普通上传"), c)
		return
	}

	// 将文件写入用户文件表
	if err := fileMetaService.UpdateUserFileMetaDB(requestData.Username , requestData.FileHash, requestData.FileName, requestData.FileSize); err != nil {
		response.BuildErrorResponse(Err.NewFastUploadError("文件写入用户文件表失败"), c)
		return
	}
	response.BuildOkResponse(http.StatusOK, "秒传成功", c)
}

func (up *UploadApi) MpUploadFileInit(c *gin.Context) {
	var requestData request.MultipleInitRequest
	if err := c.ShouldBind(&requestData); err != nil {
		log.Error(err)
		response.BuildErrorResponse(err, c)
		return
	}
	// 处理信息存入缓存
	uploadService.InitMultipleUploadFile(requestData)
	response.BuildOkResponse(http.StatusOK, "分块初始化成功", c)
}

func (up *UploadApi) UploadPart(c *gin.Context) {
	var requestData request.UploadMultipleRequest
	if err := c.ShouldBind(&requestData); err != nil {
		log.Error(err)
		response.BuildErrorResponse(err, c)
		return
	}
	err := uploadService.MultipleUploadFile(requestData)
	if err != nil {
		log.Error("multiple upload file failed err:", err)
		response.BuildErrorResponse(Err.NewMultipleUploadError("分块上传失败"), c)
	}
	response.BuildOkResponse(http.StatusOK, "分块上传成功", c)
}
