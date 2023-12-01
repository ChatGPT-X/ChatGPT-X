package user

import (
	"chatgpt_x/pkg/model"
)

// Create 创建用户，通过 User.ID 来判断是否创建成功.
func (m *User) Create() (err error) {
	if err = model.DB.Create(&m).Error; err != nil {
		return err
	}
	return nil
}

// Update 更新用户资料.
func (m *User) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&m)
	if err = model.DB.Error; err != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}

// Get 根据 ID 获取用户信息.
func Get(id int) (User, error) {
	var user User
	if err := model.DB.First(&user, id).Error; err != nil {
		return User{}, err
	}
	return User{}, nil
}
