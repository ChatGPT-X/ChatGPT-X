package ai_model

import (
	"chatgpt_x/pkg/model"
	paginator "github.com/yafeng-Soong/gorm-paginator"
)

// Create 创建 AI 模型。
func (m *AiModel) Create() (err error) {
	if err = model.DB.Create(&m).Error; err != nil {
		return err
	}
	return nil
}

// Update 更新 AI 模型。
func (m *AiModel) Update() (rowsAffected int64, err error) {
	result := model.DB.Select("*").Updates(&m)
	if err = model.DB.Error; err != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}

// Delete 删除 AI 模型。
func (m *AiModel) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(&m)
	if err = result.Error; err != nil {
		return 0, err
	}
	return result.RowsAffected, nil
}

// List 查询 AI 模型列表。
func (m *AiModel) List(aiModelType uint, page, pageSize int64) (any, error) {
	db := model.DB
	if aiModelType != 0 {
		db = db.Where("type = ?", aiModelType)
	}
	p := paginator.Page[AiModel]{
		CurrentPage: page,
		PageSize:    pageSize,
	}
	err := p.SelectPages(db)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Get 根据 ID 获取 AI 模型信息。
func Get(id uint) (AiModel, error) {
	var aiModel AiModel
	if err := model.DB.First(&aiModel, id).Error; err != nil {
		return AiModel{}, err
	}
	return aiModel, nil
}

// HasAiModelExist 通过 AiName 判断 AI 模型是否存在，存在返回 true，不存在返回 false。
func HasAiModelExist(name string, excludeID int) bool {
	var aiModel AiModel
	var count int64
	db := model.DB.Model(aiModel).Where("name = ?", name)
	if excludeID != 0 {
		db = db.Where("id != ?", excludeID)
	}
	db.Count(&count)
	return count != 0
}
