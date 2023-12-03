package bootstrap

import (
	"chatgpt_x/pkg/auth"
	"chatgpt_x/pkg/config"
)

// SetupJWT 初始化 JWT 配置。
func SetupJWT() {
	issuer := config.GetString("app.name")
	secret := config.GetString("app.jwt_secret")
	activeTime := config.GetInt("app.jwt_active_time")
	auth.InitConfig(secret, issuer, activeTime)
}
