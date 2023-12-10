package system_setting_service

import (
	"chatgpt_x/app/models/system_setting"
	"chatgpt_x/pkg/e"
)

// SystemSettingService 系统设置服务。
type SystemSettingService struct{}

// Update 更新系统设置。
func (s *SystemSettingService) Update(paramsModel system_setting.SystemSetting) (int64, e.ErrInfo) {
	systemSettingModel := system_setting.SystemSetting{
		ID:         paramsModel.ID,
		ApiBaseUrl: paramsModel.ApiBaseUrl,
		ApiProxy:   paramsModel.ApiProxy,
		ApiTimeout: paramsModel.ApiTimeout,
		WebBaseUrl: paramsModel.WebBaseUrl,
		WebProxy:   paramsModel.WebProxy,
		WebTimeout: paramsModel.WebTimeout,
	}
	rows, err := systemSettingModel.Update()
	if err != nil {
		return 0, e.ErrInfo{
			Code: e.ErrorSystemSettingSelectDetailFail,
			Msg:  err,
		}
	}
	return rows, e.ErrInfo{Code: e.SUCCESS}
}

// Detail 查询系统设置详情。
func (s *SystemSettingService) Detail() (system_setting.SystemSetting, e.ErrInfo) {
	systemSettingModel, err := system_setting.GetDetail()
	if err != nil {
		return system_setting.SystemSetting{}, e.ErrInfo{
			Code: e.ErrorSystemSettingSelectDetailFail,
			Msg:  err,
		}
	}
	return systemSettingModel, e.ErrInfo{Code: e.SUCCESS}
}
