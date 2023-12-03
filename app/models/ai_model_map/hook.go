package ai_model_map

import (
	"gorm.io/gorm"
)

// BeforeSave GORM 的钩子，在保存和更新模型前调用。
func (m *AiModelMap) BeforeSave(tx *gorm.DB) error {
	tx.Statement.Omit("create_time", "update_time")
	return nil
}
