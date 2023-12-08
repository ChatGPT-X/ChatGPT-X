package system_setting

import "chatgpt_x/pkg/model"

// GetDetail 获取系统设置详情。
func GetDetail() (SystemSetting, error) {
	var systemSetting SystemSetting
	if err := model.DB.First(&systemSetting).Error; err != nil {
		return SystemSetting{}, err
	}
	return systemSetting, nil
}

// Update 更新系统设置。
func (m *SystemSetting) Update() (rowsAffected int64, err error) {
	result := model.DB.Select("*").Updates(&m)
	if err = model.DB.Error; err != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}
