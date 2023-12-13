package setting_service

import (
	"chatgpt_x/app/models/setting"
	"chatgpt_x/app/service"
	"chatgpt_x/pkg/e"
	rds "chatgpt_x/pkg/redis"
	"context"
)

// SettingService 系统设置服务。
type SettingService struct {
	service.Service
}

var ctx = context.Background()

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

// LoadSettingsToRedis 加载系统设置到 Redis。
func (s *SettingService) LoadSettingsToRedis() e.ErrInfo {
	// 获取系统配置
	settingModel, err := setting.GetDetail()
	if err != nil {
		return e.ErrInfo{
			Code: e.ErrorSettingSelectDetailFail,
			Msg:  err,
		}
	}
	rdb := rds.RDB
	rdb.MSet(ctx, map[string]interface{}{
		service.RedisSettingOpenaiWebBaseUrl: settingModel.WebBaseUrl,
		service.RedisSettingOpenaiWebTimeout: settingModel.WebTimeout,
		service.RedisSettingOpenaiWebProxy:   settingModel.WebProxy,
		service.RedisSettingOpenaiApiBaseUrl: settingModel.ApiBaseUrl,
		service.RedisSettingOpenaiApiTimeout: settingModel.ApiTimeout,
		service.RedisSettingOpenaiApiProxy:   settingModel.ApiProxy,
	})
	return e.ErrInfo{
		Code: e.SUCCESS,
		Msg:  nil,
	}
}
