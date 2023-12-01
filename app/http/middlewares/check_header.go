package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CheckHeader 校验请求头.
func CheckHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerWk := c.GetHeader("wk")
		if headerWk != "&zq=!dL19<^2F@4PCxqspi,C" {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.Next()
	}
}
