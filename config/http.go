package config

import "chatgpt_x/pkg/config"

func init() {
	config.Add("http", config.StrMap{
		// 地址
		"listen_host": config.Env("HTTP_LISTEN_HOST", "127.0.0.1"),
		// 端口
		"listen_port": config.Env("HTTP_LISTEN_PORT", "8080"),
		// 读取包括请求体的整个请求的最大时长
		"read_timeout": config.Env("HTTP_READ_TIMEOUT", 5),
		// 允许读请求头的最大时长
		"read_header_timeout": config.Env("HTTP_READ_HEADER_TIMEOUT", 2),
		// 写响应允许的最大时长
		"write_timeout": config.Env("HTTP_WRITE_TIMEOUT", 5),
		// 当开启了保持活动状态（keep-alive）时允许的最大空闲时间
		"idle_timeout": config.Env("HTTP_IDLE_TIMEOUT", 30),
	})
}
