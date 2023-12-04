package ai_model_map

import "time"

// AI 模型关系映射表
type AiModelMap struct {
	ID          uint      `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	DisplayName string    `gorm:"column:display_name;type:varchar(100);NOT NULL" json:"display_name"`                      // 对外展示的模型名称
	AiName      string    `gorm:"column:ai_name;type:varchar(100);NOT NULL" json:"ai_name"`                                // 要映射的模型名称
	Status      uint      `gorm:"column:status;type:tinyint(4) unsigned;default:1;NOT NULL" json:"status"`                 // 状态:0-禁用 1-启用
	UpdateTime  time.Time `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"` // 修改时间
	CreateTime  time.Time `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"` // 创建时间
}

func (m *AiModelMap) TableName() string {
	return "t_ai_model_map"
}

const (
	StatusEnable  = 1 // 启用
	StatusDisable = 0 // 禁用
)
