package user

import (
	"chatgpt_x/pkg/password"
	"gorm.io/gorm"
)

// BeforeSave GORM 的钩子，在保存和更新模型前调用。
func (m *Users) BeforeSave(tx *gorm.DB) error {
	tx.Statement.Omit("create_time", "update_time")
	if !password.IsHashed(m.Password) {
		m.Password = password.Hash(m.Password)
	}
	return nil
}
