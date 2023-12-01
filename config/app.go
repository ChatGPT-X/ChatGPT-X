package config

import "chatgpt_x/pkg/config"

func init() {
	config.Add("app", config.StrMap{
		// 应用名称
		"name": config.Env("APP_NAME", "ChatGPT-X"),
		// 当前环境，用以区分多环境
		"env": config.Env("APP_ENV", "pro"),
		// APP 安全密钥，务必去创建一个自己的 GUID 作为密钥：https://www.guidgen.com
		"key": config.Env("APP_KEY", "5f90f2f4-7545-4b3e-9b96-4bcbc0343f56"),
		// 是否开启调试模式
		"debug": config.Env("APP_DEBUG", false),
		// 请求代理，为空视为不使用代理
		"fetch_proxy": config.Env("APP_FETCH_PROXY", ""),
	})
}
