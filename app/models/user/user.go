package user

import "time"

// 帐户表
type User struct {
	ID            uint      `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	AiTokenID     *uint     `gorm:"column:ai_token_id;type:int(11) unsigned" json:"ai_token_id"`                                       // 使用的密钥 id
	Username      string    `gorm:"column:username;type:varchar(50);NOT NULL" json:"username"`                                         // 账号
	Password      string    `gorm:"column:password;type:varchar(255);NOT NULL" json:"password"`                                        // 密码
	Email         string    `gorm:"column:email;type:varchar(100);NOT NULL" json:"email"`                                              // 邮箱
	IsAdmin       string    `gorm:"column:is_admin;type:enum('y','n');default:n;NOT NULL" json:"is_admin"`                             // 是否管理员：y-管理员 n-普通用户
	Status        string    `gorm:"column:status;type:enum('y','n');default:y;NOT NULL" json:"status"`                                 // 状态：y-启用 n-禁用
	LastLoginTime time.Time `gorm:"column:last_login_time;type:timestamp;default:1970-01-01 08:00:01;NOT NULL" json:"last_login_time"` // 最后登录时间
	UpdateTime    time.Time `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"`           // 修改时间
	CreateTime    time.Time `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"`           // 创建时间
}

func (m *User) TableName() string {
	return "t_users"
}

const (
	IsAdmin       = "y" // 管理员
	IsNotAdmin    = "n" // 普通用户
	StatusEnable  = "y" // 启用
	StatusDisable = "n" // 禁用
)
