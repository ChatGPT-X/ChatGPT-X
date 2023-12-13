package setting_service

import (
	"chatgpt_x/app/models/setting"
	"chatgpt_x/pkg/e"
)

// SettingService 系统设置服务。
type SettingService struct{}

// Update 更新系统设置。
func (s *SettingService) Update(paramsModel setting.Setting) (int64, e.ErrInfo) {
	settingModel := setting.Setting{
		ID:         paramsModel.ID,
		ApiBaseUrl: paramsModel.ApiBaseUrl,
		ApiProxy:   paramsModel.ApiProxy,
		ApiTimeout: paramsModel.ApiTimeout,
		WebBaseUrl: paramsModel.WebBaseUrl,
		WebProxy:   paramsModel.WebProxy,
		WebTimeout: paramsModel.WebTimeout,
	}
	rows, err := settingModel.Update()
	if err != nil {
		return 0, e.ErrInfo{
			Code: e.ErrorSettingSelectDetailFail,
			Msg:  err,
		}
	}
	return rows, e.ErrInfo{Code: e.SUCCESS}
}

// Detail 查询系统设置详情。
func (s *SettingService) Detail() (setting.Setting, e.ErrInfo) {
	settingModel, err := setting.GetDetail()
	if err != nil {
		return setting.Setting{}, e.ErrInfo{
			Code: e.ErrorSettingSelectDetailFail,
			Msg:  err,
		}
	}
	return settingModel, e.ErrInfo{Code: e.SUCCESS}
}
