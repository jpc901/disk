package middleware

import (
	"disk-master/model/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jpc901/disk-common/jwt"
)

func JwtAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			token = c.Query("token")
		}
		if token == "" {
			response.BuildOkResponse(http.StatusUnauthorized, "没有权限", c)
			c.Abort()
			return
		}
		uid, deadline, err := jwt.ParseToken(token)
		if err != nil {
			response.BuildOkResponse(http.StatusUnauthorized, "没有权限", c)
			c.Abort()
			return
		}
		if deadline < time.Now().Unix() {
			response.BuildOkResponse(http.StatusUnauthorized, "token失效，没有权限", c)
			c.Abort()
			return
		}
		c.Set("uid", uid)
		c.Next()
	}
}