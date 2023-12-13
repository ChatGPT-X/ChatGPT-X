package service

type Service struct{}

const (
	RedisSettingOpenaiWebBaseUrl = "system:openai:web:base_url"
	RedisSettingOpenaiWebTimeout = "system:openai:web:timeout"
	RedisSettingOpenaiWebProxy   = "system:openai:web:proxy"
	RedisSettingOpenaiApiBaseUrl = "system:openai:api:base_url"
	RedisSettingOpenaiApiTimeout = "system:openai:api:timeout"
	RedisSettingOpenaiApiProxy   = "system:openai:api:proxy"
)
