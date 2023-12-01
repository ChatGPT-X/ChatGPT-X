package bootstrap

import (
	"chatgpt_x/pkg/config"
	"chatgpt_x/pkg/logger"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strings"
)

// SetupLogger 初始化日志.
func SetupLogger() {
	filePath := config.GetString("logger.path")
	filePath = strings.TrimRight(filePath, "/")

	// 初始化 http请求日志
	f, _ := os.Create(filePath + "/server_http.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// 初始化 输出日志
	logger.InitZapLogger(filePath + "/server_out.log")
}
