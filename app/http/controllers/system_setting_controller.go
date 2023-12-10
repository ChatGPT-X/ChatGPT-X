package controllers

import (
	"chatgpt_x/app/models/system_setting"
	"chatgpt_x/app/requests"
	"chatgpt_x/app/service/system_setting_service"
	"chatgpt_x/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SystemSettingController 系统设置控制器。
type SystemSettingController struct {
	BaseController
}

// Update 更新系统设置。
func (ss *SystemSettingController) Update(c *gin.Context) {
	appG := ss.GetAppG(c)
	// 表单验证
	var params requests.ValidateSystemSettingUpdate
	if err := c.ShouldBind(&params); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 更新系统设置
	systemSettingService := system_setting_service.SystemSettingService{}
	rows, errInfo := systemSettingService.Update(system_setting.SystemSetting{
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
	appG.Response(http.StatusOK, e.SUCCESS, nil, gin.H{"rows": rows})
}

// Detail 查询系统设置详情。
func (ss *SystemSettingController) Detail(c *gin.Context) {
	appG := ss.GetAppG(c)
	// 查询系统设置详情
	systemSettingService := system_setting_service.SystemSettingService{}
	systemSettingModel, errInfo := systemSettingService.Detail()
	if errInfo.Code != e.SUCCESS {
		appG.Response(http.StatusOK, errInfo.Code, errInfo.Msg, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, systemSettingModel)
}
