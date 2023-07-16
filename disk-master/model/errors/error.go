package errors

import "fmt"

type Error struct {
	Code    ErrorCode
	Message string
	Detail  interface{}
}

func (e *Error) Error() string {
	return fmt.Sprintf("[DiskError] code: %d, msg: %s", e.Code, e.Message)
}

func NewInternalServerError(detail interface{}) *Error {
	return &Error{
		Code:    InternalServerErrorCode,
		Message: "internal server error",
		Detail:  detail,
	}
}

func NewUploadFailedError(detail interface{}) *Error {
	return &Error{
		Code:    FileUploadFailedCode,
		Message: "upload failed",
		Detail:  detail,
	}
}

func NewGetFileMetaError(detail interface{}) *Error {
	return &Error{
		Code: GetFileMetaFailedCode,
		Message: "get file meta failed",
		Detail: detail,
	}
}
func NewFileUpdateError(detail interface{}) *Error {
	return &Error{
		Code: FileUpdateFailedCode,
		Message: "file meta update failed",
		Detail: detail,
	}
}

func NewUserSignUpError(detail interface{}) *Error {
	return &Error{
		Code: UserSignUpFailedCode,
		Message: "user sign up failed",
		Detail: detail,
	}
}

func NewUserExistError(detail interface{}) *Error {
	return &Error{
		Code: UserExistCode,
		Message: "用户已存在",
		Detail: detail,
	}
}

func NewUserNotExistError(detail interface{}) *Error {
	return &Error{
		Code: UserNotExistCode,
		Message: "用户不存在",
		Detail: detail,
	}
}

func NewPasswordError(detail interface{}) *Error {
	return &Error{
		Code: PasswordErrorCode,
		Message: "密码错误",
		Detail: detail,
	}
}

