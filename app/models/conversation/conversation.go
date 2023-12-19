package conversation

import "time"

// 对话表
type Conversation struct {
	ID                uint      `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	UserID            uint      `gorm:"column:user_id;type:int(11) unsigned;NOT NULL" json:"user_id"`                            // 用户 id
	AiTokenID         *uint     `gorm:"column:ai_token_id;type:int(11) unsigned" json:"ai_token_id"`                             // AI 密钥 id
	Type              string    `gorm:"column:type;type:enum('web','api');default:web;NOT NULL" json:"type"`                     // 对话类型：web api
	ModelName         string    `gorm:"column:model_name;type:varchar(255);NOT NULL" json:"model_name"`                          // 模型名称
	ConversationID    string    `gorm:"column:conversation_id;type:varchar(36);NOT NULL" json:"conversation_id"`                 // 对话 id
	ConversationTitle string    `gorm:"column:conversation_title;type:varchar(128);NOT NULL" json:"conversation_title"`          // 对话标题
	Status            string    `gorm:"column:status;type:enum('y','n');default:y;NOT NULL" json:"status"`                       // 状态：y-正常 n-删除
	UpdateTime        time.Time `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"` // 修改时间
	CreateTime        time.Time `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"` // 创建时间
}

func (m *Conversation) TableName() string {
	return "t_conversation"
}

const (
	TypeWeb       = "web" // WEB
	TypeApi       = "api" // API
	StatusEnable  = "y"
	StatusDisable = "n"
)
