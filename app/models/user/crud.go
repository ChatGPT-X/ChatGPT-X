package user

import (
	"chatgpt_x/pkg/model"
	"chatgpt_x/pkg/password"
)

// Create 创建用户，通过 User.ID 来判断是否创建成功。
func (m *Users) Create() (err error) {
	if err = model.DB.Create(&m).Error; err != nil {
		return err
	}
	return nil
}

// Update 更新用户资料。
func (m *Users) Update() (rowsAffected int64, err error) {
	result := model.DB.Select("*").Updates(&m)
	if err = model.DB.Error; err != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}

// IsDisable 判断用户是否被禁用，禁用返回 true，未禁用返回 false。
func (m *Users) IsDisable() bool {
	return m.IsDisabled == IsDisabled
}

// Get 根据 ID 获取用户信息。
func Get(id int) (Users, error) {
	var user Users
	if err := model.DB.First(&user, id).Error; err != nil {
		return Users{}, err
	}
	return Users{}, nil
}

// HasByUsernameExist 通过 Username 判断用户是否存在，存在返回 true，不存在返回 false。
func HasByUsernameExist(username string) bool {
	var user Users
	var count int64
	model.DB.Model(user).Where("username = ?", username).Count(&count)
	return count != 0
}

// GetByUsername 通过 Username 获取用户信息。
func GetByUsername(username string) (Users, error) {
	var user Users
	if err := model.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return Users{}, err
	}
	return user, nil
}

// CheckPassword 检查密码是否正确, 正确返回 true，错误返回 false。
func CheckPassword(pass, hash string) bool {
	return password.IsHashed(hash) && password.CheckHash(pass, hash)
}
