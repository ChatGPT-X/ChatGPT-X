package bootstrap

import (
	"chatgpt_x/app/models/setting"
	"chatgpt_x/app/service/openai_service"
	"fmt"
)

// SetupOpenAI 初始化 OpenAI 配置。
func SetupOpenAI() {
	// 获取系统配置
	settingModel, err := setting.GetDetail()
	if err != nil {
		panic(fmt.Sprintf("Init openai web config fail: %s", err))
	}
	openai_service.Setting = map[string]any{
		"WebBaseUrl": settingModel.WebBaseUrl,
		"WebTimeout": settingModel.WebTimeout,
		"ApiBaseUrl": settingModel.ApiBaseUrl,
		"ApiTimeout": settingModel.ApiTimeout,
	}
}
