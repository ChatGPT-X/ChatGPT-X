package bootstrap

import (
	"chatgpt_x/app/models/system_setting"
	"chatgpt_x/app/service/openai_service"
	"fmt"
)

func SetupOpenAI() {
	// 获取系统配置
	systemSetting, err := system_setting.GetDetail()
	if err != nil {
		panic(fmt.Sprintf("Init openai web config fail: %s", err))
	}
	openai_service.WebBaseUrl = systemSetting.WebBaseUrl
	openai_service.ApiBaseUrl = systemSetting.ApiBaseUrl
}
