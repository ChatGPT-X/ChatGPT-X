package ai_model_service

import (
	"chatgpt_x/app/models/ai_model"
	"chatgpt_x/pkg/e"
	"fmt"
	paginator "github.com/yafeng-Soong/gorm-paginator"
)

// AiModelService AI 模型服务。
type AiModelService struct{}

// Create 创建 AI 模型。
func (s *AiModelService) Create(paramsModel ai_model.AiModel) e.ErrInfo {
	// 检查 AI 模型是否存在
	if ai_model.HasAiModelExist(paramsModel.Name, 0) {
		return e.ErrInfo{
			Code: e.ErrorAiModelIsExist,
			Msg:  fmt.Errorf("ai model is exist: %s", paramsModel.Name),
		}
	}
	// 创建 AI 模型
	aiModel := ai_model.AiModel{
		Type:      paramsModel.Type,
		AliasName: paramsModel.AliasName,
		Name:      paramsModel.Name,
		Status:    paramsModel.Status,
	}
	if err := aiModel.Create(); err != nil {
		return e.ErrInfo{
			Code: e.ErrorAiModelCreateFail,
			Msg:  err,
		}
	}
	return e.ErrInfo{Code: e.SUCCESS}
}

// Update 更新 AI 模型。
func (s *AiModelService) Update(paramsModel ai_model.AiModel) (int64, e.ErrInfo) {
	// 检查 AI 模型是否存在
	if ai_model.HasAiModelExist(paramsModel.Name, int(paramsModel.ID)) {
		return 0, e.ErrInfo{
			Code: e.ErrorAiModelIsExist,
			Msg:  fmt.Errorf("ai model is exist: %s", paramsModel.Name),
		}
	}
	// 更新 AI 模型
	aiModel := ai_model.AiModel{
		ID:        paramsModel.ID,
		Type:      paramsModel.Type,
		AliasName: paramsModel.AliasName,
		Name:      paramsModel.Name,
		Status:    paramsModel.Status,
	}
	rows, err := aiModel.Update()
	if err != nil {
		return 0, e.ErrInfo{
			Code: e.ErrorAiModelUpdateFail,
			Msg:  err,
		}
	}
	return rows, e.ErrInfo{Code: e.SUCCESS}
}

// List 查询 AI 模型列表。
func (s *AiModelService) List(aiModelType uint, page, pageSize int64) (paginator.Page[ai_model.AiModel], e.ErrInfo) {
	aiModel := ai_model.AiModel{}
	pageData, err := aiModel.List(aiModelType, page, pageSize)
	if err != nil {
		return paginator.Page[ai_model.AiModel]{}, e.ErrInfo{
			Code: e.ErrorAiModelSelectListFail,
			Msg:  err,
		}
	}
	data := pageData.(paginator.Page[ai_model.AiModel])
	return data, e.ErrInfo{Code: e.SUCCESS}
}

// Delete 删除 AI 模型。
func (s *AiModelService) Delete(id uint) (int64, e.ErrInfo) {
	aiModel := ai_model.AiModel{
		ID: id,
	}
	rows, err := aiModel.Delete()
	if err != nil {
		return 0, e.ErrInfo{
			Code: e.ErrorAiModelDeleteFail,
			Msg:  err,
		}
	}
	return rows, e.ErrInfo{Code: e.SUCCESS}
}
