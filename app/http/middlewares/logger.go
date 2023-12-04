package middlewares

import (
	"chatgpt_x/pkg/app"
	"chatgpt_x/pkg/config"
	"chatgpt_x/pkg/e"
	"chatgpt_x/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime/debug"
)

var appName = config.Get("app.name")

func Logger() gin.HandlerFunc {
	return gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: logFormatter,
		Output:    gin.DefaultWriter,
		SkipPaths: nil,
	})
}

func logFormatter(param gin.LogFormatterParams) string {
	var statusColor, methodColor, resetColor string
	if param.IsOutputColor() {
		statusColor = param.StatusCodeColor()
		methodColor = param.MethodColor()
		resetColor = param.ResetColor()
	}

	return fmt.Sprintf("[%s] %v |%s %3d %s| %13v | %15s |%s %-7s %s %#v %s\n%s",
		appName,
		param.TimeStamp.Format("2006/01/02 - 15:04:05"),
		statusColor, param.StatusCode, resetColor,
		param.Latency,
		param.ClientIP,
		methodColor, param.Method, resetColor,
		param.Path,
		param.Request.Form,
		param.ErrorMessage,
	)
}

func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录错误和堆栈跟踪
				logger.Error(fmt.Errorf("Panic Recovered: %s\n%s\n", err, debug.Stack()))
				// 返回错误响应
				errCode := e.ERROR
				c.AbortWithStatusJSON(http.StatusInternalServerError, app.Response{
					Code: errCode,
					Msg:  e.GetMsg(errCode),
					Data: nil,
				})
				return
			}
		}()
		c.Next() // 继续执行后续的处理程序和中间件
	}
}
