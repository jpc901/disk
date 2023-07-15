package errors

type ErrorCode uint32

const (
	InvalidParameterCode    ErrorCode = 400
	InternalServerErrorCode ErrorCode = 500
	FileUploadFailedCode    ErrorCode = 1001
)
