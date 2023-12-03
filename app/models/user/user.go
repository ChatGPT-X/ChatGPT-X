package user

import "time"

// 帐户表
type Users struct {
	ID            uint      `gorm:"column:id;type:int(11) unsigned;primary_key;AUTO_INCREMENT" json:"id"`
	Username      string    `gorm:"column:username;type:varchar(50);NOT NULL" json:"username"`                               // 账号
	Password      string    `gorm:"column:password;type:varchar(255);NOT NULL" json:"password"`                              // 密码
	Email         string    `gorm:"column:email;type:varchar(100);NOT NULL" json:"email"`                                    // 邮箱
	IsAdmin       uint      `gorm:"column:is_admin;type:tinyint(4) unsigned;default:0;NOT NULL" json:"is_admin"`             // 是否管理员:0-普通用户 1-管理员
	IsDisabled    uint      `gorm:"column:is_disabled;type:tinyint(4) unsigned;default:0;NOT NULL" json:"is_disabled"`       // 是否禁用:0-启用 1-封禁
	LastLoginTime time.Time `gorm:"column:last_login_time;type:timestamp;NOT NULL" json:"last_login_time"`                   // 最后登录时间
	UpdateTime    time.Time `gorm:"column:update_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"` // 修改时间
	CreateTime    time.Time `gorm:"column:create_time;type:timestamp;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"` // 创建时间
}

func (m *Users) TableName() string {
	return "t_users"
}

const (
	IsAdmin       = 1 // 管理员
	IsNotAdmin    = 0 // 普通用户
	IsDisabled    = 1 // 禁用
	IsNotDisabled = 0 // 启用
)
