package response

import (
	"disk-master/model/errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Error struct {
	Message string      `json:"message"`
	Detail  interface{} `json:"detail"`
}

type ErrorResponse struct {
	RequestId string `json:"-"`
	Code      int    `json:"code"`
	Error     Error  `json:"error"`
}

type OkResponse struct {
	RequestId string      `json:"-"`
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
}

func buildErrorResult(code int, msg string, detail interface{}, c *gin.Context) {
	err := Error{
		msg,
		detail,
	}
	httpStatus := code
	if code < 100 || code > 999 {
		httpStatus = http.StatusOK
	}
	c.JSON(httpStatus, ErrorResponse{
		c.Request.Header.Get("Request-Id"),
		code,
		err,
	})
}

func BuildOkResponse(code int, data interface{}, c *gin.Context) {
	c.JSON(code, OkResponse{
		c.Request.Header.Get("Request-Id"),
		0,
		data,
	})
}

func BuildHtmlResponse(path string, c *gin.Context) {
	c.HTML(http.StatusOK, path, nil)
}

func BuildErrorResponse(e error, c *gin.Context) {
	if e == nil {
		buildErrorResult(http.StatusInternalServerError, "internal server error", "", c)
	} else if err, ok := e.(*errors.Error); ok {
		buildErrorResult(int(err.Code), err.Message, err.Detail, c)
	} else if err, ok := e.(validator.ValidationErrors); ok {
		buildErrorResult(http.StatusBadRequest, "invalid parameter", err.Error(), c)
	} else {
		buildErrorResult(http.StatusInternalServerError, "internal server error", e.Error(), c)
	}
}