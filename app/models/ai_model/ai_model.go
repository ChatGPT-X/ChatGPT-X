package ai_model

import "time"

// AI 模型表
type AiModel struct {
	ID         uint      `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Type       uint      `gorm:"column:type;type:tinyint(4) unsigned;default:1;NOT NULL" json:"type"`                     // AI 模型类型：1-api 2-web
	AliasName  string    `gorm:"column:alias_name;type:varchar(100);NOT NULL" json:"alias_name"`                          // AI 模型别名
	Name       string    `gorm:"column:name;type:varchar(100);NOT NULL" json:"name"`                                      // AI 模型完整名称
	Status     string    `gorm:"column:status;type:enum('y','n');default:y;NOT NULL" json:"status"`                       // 状态：y-启用 n-禁用
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"` // 修改时间
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"` // 创建时间
}

func (m *AiModel) TableName() string {
	return "t_ai_models"
}

const (
	TypeApi       = 1   // API
	TypeWeb       = 2   // WEB
	StatusEnable  = "y" // 启用
	StatusDisable = "n" // 禁用
)
