package bootstrap

import (
	"chatgpt_x/app/models/system_setting"
	"chatgpt_x/app/service/openai_service"
	"fmt"
)

// SetupOpenAI 初始化 OpenAI 配置。
func SetupOpenAI() {
	// 获取系统配置
	systemSetting, err := system_setting.GetDetail()
	if err != nil {
		panic(fmt.Sprintf("Init openai web config fail: %s", err))
	}
	openai_service.SystemSetting = map[string]any{
		"WebBaseUrl": systemSetting.WebBaseUrl,
		"WebTimeout": systemSetting.WebTimeout,
		"ApiBaseUrl": systemSetting.ApiBaseUrl,
		"ApiTimeout": systemSetting.ApiTimeout,
	}
}
