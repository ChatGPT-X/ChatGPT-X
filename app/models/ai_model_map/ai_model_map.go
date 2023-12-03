package ai_model_map

import "time"

// AI 大模型关系映射表
type AiModelMap struct {
	ID          uint      `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	DisplayName string    `gorm:"column:display_name;type:varchar(100);NOT NULL" json:"display_name"`                      // 对外展示的模型名称
	AiName      string    `gorm:"column:ai_name;type:varchar(100);NOT NULL" json:"ai_name"`                                // 要映射的模型名称
	IsDisabled  uint      `gorm:"column:is_disabled;type:tinyint(4) unsigned;default:0;NOT NULL" json:"is_disabled"`       // 是否禁用:0-启用 1-封禁
	UpdateTime  time.Time `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"` // 修改时间
	CreateTime  time.Time `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"` // 创建时间
}

func (m *AiModelMap) TableName() string {
	return "t_ai_model_map"
}
