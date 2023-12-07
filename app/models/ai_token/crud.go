package ai_token

import (
	"chatgpt_x/pkg/model"
	paginator "github.com/yafeng-Soong/gorm-paginator"
)

// Create 创建密钥。
func (m *AiToken) Create() (err error) {
	if err = model.DB.Create(&m).Error; err != nil {
		return err
	}
	return nil
}

// Update 更新密钥。
func (m *AiToken) Update() (rowsAffected int64, err error) {
	result := model.DB.Select("*").Updates(&m)
	if err = model.DB.Error; err != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}

// Delete 删除密钥。
func (m *AiToken) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(&m)
	if err = result.Error; err != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}

// List 查询密钥列表。
func (m *AiToken) List(page, pageSize int64) (any, error) {
	db := model.DB
	p := paginator.Page[AiToken]{
		CurrentPage: page,
		PageSize:    pageSize,
	}
	err := p.SelectPages(db)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// HasTokenExist 判断密钥是否存在，存在返回 true，不存在返回 false。
func HasTokenExist(token string, excludeID int) bool {
	var aiTokenModel AiToken
	var count int64
	db := model.DB.Model(aiTokenModel).Where("token = ?", token)
	if excludeID != 0 {
		db = db.Where("id != ?", excludeID)
	}
	db.Count(&count)
	return count != 0
}
