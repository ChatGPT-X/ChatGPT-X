package bootstrap

import (
	"chatgpt_x/config"
	"github.com/gin-gonic/gin"
)

// Setup 初始化指定的服务.
func Setup() {
	gin.SetMode(getMode())
	autoLoader(
		config.Initialize, // Configs initialize file.
		SetupJWT,          // JWT
		SetupMySQL,        // MySQL
		SetupRedis,        // Redis
		SetupLogger,       // Logs
	)
}

// autoLoader 自动加载初始化.
func autoLoader(funcName ...func()) {
	// 只是单纯的初始化服务模块，没有参数，没有返回值！！
	for _, v := range funcName {
		v()
	}
}
