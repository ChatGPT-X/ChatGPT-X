package token

import "time"

// 密钥表
type Token struct {
	ID         uint      `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Type       uint      `gorm:"column:type;type:tinyint(4) unsigned;default:1;NOT NULL" json:"type"`                     // 类型:0-apiKey 1-accessToken
	Token      string    `gorm:"column:token;type:text;NOT NULL" json:"token"`                                            // 密钥内容
	Remark     string    `gorm:"column:remark;type:varchar(255);NOT NULL" json:"remark"`                                  // 备注
	Status     uint      `gorm:"column:status;type:tinyint(4) unsigned;default:1;NOT NULL" json:"status"`                 // 状态:0-禁用 1-启用
	UpdateTime time.Time `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"` // 修改时间
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"` // 创建时间
}

func (m *Token) TableName() string {
	return "t_tokens"
}

const (
	TypeApiKey      = 0 // apiKey
	TypeAccessToken = 1 // accessToken
	StatusDisable   = 0 // 禁用
	StatusEnable    = 1 // 启用
)
