package controllers

import (
	"chatgpt_x/app/models/setting"
	"chatgpt_x/app/requests"
	"chatgpt_x/app/service/setting_service"
	"chatgpt_x/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SettingController 系统设置控制器。
type SettingController struct {
	BaseController
}

// Update 更新系统设置。
func (s *SettingController) Update(c *gin.Context) {
	appG := s.GetAppG(c)
	// 表单验证
	var params requests.ValidateSettingUpdate
	if err := c.ShouldBind(&params); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 更新系统设置
	settingService := setting_service.SettingService{}
	rows, errInfo := settingService.Update(setting.Setting{
		ID:         params.ID,
		ApiBaseUrl: params.ApiBaseUrl,
		ApiProxy:   params.ApiProxy,
		ApiTimeout: params.ApiTimeout,
		WebBaseUrl: params.WebBaseUrl,
		WebProxy:   params.WebProxy,
		WebTimeout: params.WebTimeout,
	})
	if errInfo.Code != e.SUCCESS {
		appG.Response(http.StatusOK, errInfo.Code, errInfo.Msg, nil)
		return
	}
	// 重新从将配置信息读取到 Redis 使用
	err := settingService.LoadSettingsToRedis()
	if err.Code != e.SUCCESS {
		appG.Response(http.StatusOK, err.Code, err.Msg, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, gin.H{"rows": rows})
}

// Detail 获取系统设置详情。
func (s *SettingController) Detail(c *gin.Context) {
	appG := s.GetAppG(c)
	// 获取系统设置详情
	settingService := setting_service.SettingService{}
	settingModel, errInfo := settingService.Detail()
	if errInfo.Code != e.SUCCESS {
		appG.Response(http.StatusOK, errInfo.Code, errInfo.Msg, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, settingModel)
}
