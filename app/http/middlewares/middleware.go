package middlewares

import (
	"chatgpt_x/pkg/config"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
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
		RequestID(),                          // 为每个请求标记一个唯一性质的 ID
		Logger(),                             // 访问请求记录到日志
		CustomRecovery(),                     // 恢复从 panic 中间件
		sessions.Sessions("SESSIONS", store), // SESSION
	)
}
