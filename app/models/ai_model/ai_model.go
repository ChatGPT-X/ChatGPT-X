package ai_model

import "time"

// AI 模型表
type AiModel struct {
	ID         uint      `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Type       string    `gorm:"column:type;type:enum('web','api');default:web;NOT NULL" json:"type"`                     // 类型：web api
	AliasName  string    `gorm:"column:alias_name;type:varchar(100);NOT NULL" json:"alias_name"`                          // 别名
	Name       string    `gorm:"column:name;type:varchar(100);NOT NULL" json:"name"`                                      // 完整名称
	Status     string    `gorm:"column:status;type:enum('y','n');default:y;NOT NULL" json:"status"`                       // 状态：y-启用 n-禁用
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"` // 修改时间
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"` // 创建时间
}

func (m *AiModel) TableName() string {
	return "t_ai_models"
}

const (
	TypeWeb       = "web" // WEB
	TypeApi       = "api" // API
	StatusEnable  = "y"   // 启用
	StatusDisable = "n"   // 禁用
)
