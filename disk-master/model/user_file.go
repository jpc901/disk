package model

type UserFileMeta struct {
	UserName    string 	`json:"username"`
	FileHash    string	`json:"fileHash"`
	FileName    string	`json:"fileName"`
	FileSize    int64	`json:"fileSize"`
	UploadAt    string	`json:"uploadAt"`
	LastUpdated string  `json:"lastUpdated"`
}