package health

import "github.com/gin-gonic/gin"

type HealthyApi struct{}

func (h *HealthyApi) Liveness(c *gin.Context) {
	c.JSON(200, "ok")
}

func (h *HealthyApi) Readiness(c *gin.Context) {
	c.JSON(200, "ok")
}
