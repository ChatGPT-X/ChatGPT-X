package ai_token

import "time"

// AI 密钥表
type AiToken struct {
	ID         uint      `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Type       string    `gorm:"column:type;type:enum('access_token','api_key');default:access_token;NOT NULL" json:"type"` // 类型：access_token api_key
	Token      string    `gorm:"column:token;type:text;NOT NULL" json:"token"`                                              // 密钥内容
	Remark     string    `gorm:"column:remark;type:varchar(255);NOT NULL" json:"remark"`                                    // 备注
	Status     string    `gorm:"column:status;type:enum('y','n');default:y;NOT NULL" json:"status"`                         // 状态：y-启用 n-禁用
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"`   // 修改时间
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"`   // 创建时间
}

func (m *AiToken) TableName() string {
	return "t_ai_tokens"
}

const (
	TypeAccessToken = "access_token" // accessToken
	TypeApiKey      = "api_key"      // apiKey
	StatusEnable    = "y"            // 启用
	StatusDisable   = "n"            // 禁用
)
