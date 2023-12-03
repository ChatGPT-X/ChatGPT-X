package config

import "chatgpt_x/pkg/config"

func init() {
	config.Add("app", config.StrMap{
		// 应用名称
		"name": config.Env("APP_NAME", "ChatGPT-X"),
		// 当前环境，用以区分多环境
		"env": config.Env("APP_ENV", "pro"),
		// JWT 安全密钥，项目初始化前，务必去创建一个64位的随机字符串做自己的的密钥
		"jwt_secret": config.Env("APP_JWT_SECRET", "sn7xtwpoimzer8cu641gy9v0qjfk2labh53drlgucrllpusn7h8hu6ruje419x2k"),
		// JWT 过期时间（单位：秒），推荐 7200 秒，即两个小时后过期
		"jwt_active_time": config.Env("APP_JWT_ACTIVE_TIME", 7200),
		// 是否开启调试模式
		"debug": config.Env("APP_DEBUG", false),
		// 请求代理，为空视为不使用代理
		"fetch_proxy": config.Env("APP_FETCH_PROXY", ""),
	})
}
