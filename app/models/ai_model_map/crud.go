package ai_model_map

import "chatgpt_x/pkg/model"

// Create 创建 AI 模型关系映射。
func (m *AiModelMap) Create() (err error) {
	if err = model.DB.Create(&m).Error; err != nil {
		return err
	}
	return nil
}

// Update 更新 AI 模型关系映射。
func (m *AiModelMap) Update() (rowsAffected int64, err error) {
	result := model.DB.Select("*").Updates(&m)
	if err = model.DB.Error; err != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}

// Delete 删除 AI 模型关系映射。
func (m *AiModelMap) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(&m)
	if err = result.Error; err != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}

// HasAiModelExist 通过 AiName 判断AI模型是否存在，存在返回 true，不存在返回 false。
func HasAiModelExist(aiName string, excludeID int) bool {
	var aiModelMap AiModelMap
	var count int64
	db := model.DB.Model(aiModelMap).Where("ai_name = ?", aiName)
	if excludeID != 0 {
		db = db.Where("id != ?", excludeID)
	}
	db.Count(&count)
	return count != 0
}
