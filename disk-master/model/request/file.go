package request

import "mime/multipart"

type UploadFileRequest struct {
	File     *multipart.FileHeader `form:"file" binding:"required"`
}
type FileMetaRequest struct {
	FileHash string `form:"fileHash" binding:"required"`
}

type FileDownloadRequest struct {
	FileHash string `form:"fileHash" binding:"required"`
}

type FileUpdateRequest struct {
	FileHash    string `form:"fileHash" binding:"required"`
	FileName    string `form:"fileName" binding:"required"`
	OperateType string `form:"operateType" binding:"required"`
}

type FileDeleteRequest struct {
	FileHash string `json:"fileHash" binding:"required"`
	FileName string `json:"fileName" binding:"required"`
}

type FastUploadRequest struct {
	Username string `form:"username" binding:"required"`
	FileHash string `form:"fileHash" binding:"required"`
	FileName string `form:"fileName" binding:"required"`
}

type MultipleInitRequest struct {
	UploadId   string `json:"uploadId" binding:"required"`
	FileHash   string `json:"fileHash" binding:"required"`
	FileName   string `json:"fileName" binding:"required"`
	FileSize   int64  `json:"fileSize" binding:"required"`
	ChunkCount int64  `json:"chunkCount" binding:"omitempty"`
	ChunkSize  int64  `json:"chunkSize" binding:"omitempty"`
}

type UploadMultipleRequest struct {
	File      *multipart.FileHeader `form:"file" binding:"required"`
	UploadId  string                `form:"uploadId" binding:"required"`
	FileHash  string                `form:"fileHash" binding:"required"`
	FileName  string                `form:"fileName" binding:"required"`
	FileSize  int64                 `form:"fileSize" binding:"required"`
	CurChunk  string                `form:"curChunk" binding:"required"`
	ChunkSize int64                 `form:"chunkSize" binding:"required"`
}

type CheckChunkExistRequest struct {
	UploadId string `form:"uploadId" binding:"required"`
	CurChunk string `form:"curChunk" binding:"required"`
}
