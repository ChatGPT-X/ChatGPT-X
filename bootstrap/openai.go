package bootstrap

import (
	"chatgpt_x/app/service/setting_service"
	"chatgpt_x/pkg/e"
	"fmt"
)

// SetupOpenAI 初始化 OpenAI 配置。
func SetupOpenAI() {
	// 获取系统配置
	settingService := setting_service.SettingService{}
	errInfo := settingService.LoadSettingsToRedis()
	if errInfo.Code != e.SUCCESS {
		panic(fmt.Sprintf("Init openai web config fail: %s", errInfo.Msg))
	}
}
