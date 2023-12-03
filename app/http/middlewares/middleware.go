package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Register used for register middleware.
func Register(engine *gin.Engine) {
	engine.Use(
		Logger(),         // 访问请求记录到日志
		CustomRecovery(), // 恢复从 panic 中间件
	)
}
