package request

import "mime/multipart"

type UploadFileRequest struct{
	File *multipart.FileHeader `form:"file" binding:"required"`
}

type FileMetaRequest struct {
	FileHash string `form:"fileHash" binding:"required"`
}

type FileDownloadRequest struct {
	FileHash string `form:"fileHash" binding:"required"`
}

type FileUpdateRequest struct {
	FileHash string `form:"fileHash" binding:"required"`
	FileName string `form:"fileName" binding:"required"`
	OperateType string `form:"operateType" binding:"required"`
}

type FileDeleteRequest struct {
	FileHash string `form:"fileHash" binding:"required"`
}