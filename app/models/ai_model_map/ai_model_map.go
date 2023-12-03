package ai_model_map

import (
	"time"
)

// AI 模型关系映射表
type AiModelMap struct {
	ID          uint      `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	DisplayName string    `gorm:"column:display_name;NOT NULL"`                          // 对外展示的模型名称
	AiName      string    `gorm:"column:ai_name;NOT NULL"`                               // 要映射的模型名称
	IsDisabled  uint      `gorm:"column:is_disabled;default:0;NOT NULL"`                 // 是否禁用:0-启用 1-封禁
	UpdateTime  time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 修改时间
	CreateTime  time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
}

func (m *AiModelMap) TableName() string {
	return "t_ai_model_map"
}
