package errors

type ErrorCode uint32

const (
	InvalidParameterCode    ErrorCode = 400
	InternalServerErrorCode ErrorCode = 500
	FileUploadFailedCode    ErrorCode = 1001
	GetFileMetaFailedCode   ErrorCode = 1002
	FileUpdateFailedCode    ErrorCode = 1003
	UserSignUpFailedCode    ErrorCode = 1004
	UserExistCode			ErrorCode = 1005
	UserNotExistCode 		ErrorCode = 1006
	PasswordErrorCode		ErrorCode = 1007
	GetUserFileErrorCode 	ErrorCode = 1008
	FastUploadFileErrorCode	ErrorCode = 1009
	MultipleInitFailedCode  ErrorCode = 1010
)
