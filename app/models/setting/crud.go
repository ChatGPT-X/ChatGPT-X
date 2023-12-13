package setting

import "chatgpt_x/pkg/model"

// GetDetail 获取系统设置详情。
func GetDetail() (Setting, error) {
	var setting Setting
	if err := model.DB.First(&setting).Error; err != nil {
		return Setting{}, err
	}
	return setting, nil
}

// Update 更新系统设置。
func (m *Setting) Update() (rowsAffected int64, err error) {
	result := model.DB.Select("*").Updates(&m)
	if err = model.DB.Error; err != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}
