package user

import (
	"time"
)

// 帐户表
type Users struct {
	ID         uint      `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Username   string    `gorm:"column:username;NOT NULL"`                              // 账号
	Password   string    `gorm:"column:password;NOT NULL"`                              // 密码
	Email      string    `gorm:"column:email;NOT NULL"`                                 // 邮箱
	IsAdmin    uint      `gorm:"column:is_admin;default:0;NOT NULL"`                    // 是否管理员:0-普通用户 1-管理员
	IsDisabled uint      `gorm:"column:is_disabled;default:0;NOT NULL"`                 // 是否禁用:0-启用 1-封禁
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 修改时间
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
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
