package middlewares

import (
	"chatgpt_x/pkg/config"
	"chatgpt_x/pkg/logger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"time"
)

// Register used for register middleware.
func Register(engine *gin.Engine) {
	store := cookie.NewStore([]byte(config.GetString("app.key")))
	store.Options(sessions.Options{
		Path:     "/",
		Domain:   config.GetString("http.listen_host"),
		MaxAge:   86400 * 3,
		Secure:   false,
		HttpOnly: true,
		SameSite: 0,
	})

	engine.Use(
		RequestID(),                               // 为每个请求标记一个唯一性质的 ID
		GinRecoveryWithZap(logger.Logger, true),   // err 和 panic 记录到日志（包括堆栈信息）
		GinZap(logger.Logger, time.RFC3339, true), // 访问请求记录到日志
		sessions.Sessions("SESSIONS", store),      // SESSION
	)
}
