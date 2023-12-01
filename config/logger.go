package config

import "chatgpt_x/pkg/config"

func init() {
	config.Add("logger", config.StrMap{
		// 日志存放目录(含请求日志、输出日志等... 不需要指定日志名称，目录要提前建好)
		"path": config.Env("LOGGER_PATH", "./logs"),
	})
}
