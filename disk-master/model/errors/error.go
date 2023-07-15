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