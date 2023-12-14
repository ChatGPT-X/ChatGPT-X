package ai_token_service

import (
	"chatgpt_x/app/models/ai_token"
	"chatgpt_x/pkg/e"
	"fmt"
	paginator "github.com/yafeng-Soong/gorm-paginator"
)

// AiTokenService AI 密钥服务。
type AiTokenService struct{}

// Create 创建 AI 密钥。
func (s *AiTokenService) Create(paramsModel ai_token.AiToken) e.ErrInfo {
	// 检查密钥是否存在
	if ai_token.HasTokenExist(paramsModel.Token, 0) {
		return e.ErrInfo{
			Code: e.ErrorAiTokenIsExist,
			Msg:  fmt.Errorf("ai token is exist: %s", paramsModel.Token),
		}
	}
	// 创建密钥
	aiTokenModel := ai_token.AiToken{
		Type:   paramsModel.Type,
		Token:  paramsModel.Token,
		Remark: paramsModel.Remark,
		Status: paramsModel.Status,
	}
	if err := aiTokenModel.Create(); err != nil {
		return e.ErrInfo{
			Code: e.ErrorAiTokenCreateFail,
			Msg:  err,
		}
	}
	return e.ErrInfo{Code: e.SUCCESS}
}

// Update 更新 AI 密钥。
func (s *AiTokenService) Update(paramsModel ai_token.AiToken) (int64, e.ErrInfo) {
	// 检查密钥是否存在
	if ai_token.HasTokenExist(paramsModel.Token, int(paramsModel.ID)) {
		return 0, e.ErrInfo{
			Code: e.ErrorAiTokenIsExist,
			Msg:  fmt.Errorf("ai token is exist: %s", paramsModel.Token),
		}
	}
	// 更新密钥
	aiTokenModel := ai_token.AiToken{
		ID:     paramsModel.ID,
		Type:   paramsModel.Type,
		Token:  paramsModel.Token,
		Remark: paramsModel.Remark,
		Status: paramsModel.Status,
	}
	rows, err := aiTokenModel.Update()
	if err != nil {
		return 0, e.ErrInfo{
			Code: e.ErrorAiTokenUpdateFail,
			Msg:  err,
		}
	}
	return rows, e.ErrInfo{Code: e.SUCCESS}
}

// List 获取 AI 密钥列表。
func (s *AiTokenService) List(page, pageSize int64) (paginator.Page[ai_token.AiToken], e.ErrInfo) {
	aiTokenModel := ai_token.AiToken{}
	pageData, err := aiTokenModel.List(page, pageSize)
	if err != nil {
		return paginator.Page[ai_token.AiToken]{}, e.ErrInfo{
			Code: e.ErrorAiTokenSelectListFail,
			Msg:  err,
		}
	}
	data := pageData.(paginator.Page[ai_token.AiToken])
	return data, e.ErrInfo{Code: e.SUCCESS}
}

// Delete 删除 AI 密钥。
func (s *AiTokenService) Delete(id uint) (int64, e.ErrInfo) {
	aiTokenModel := ai_token.AiToken{
		ID: id,
	}
	rows, err := aiTokenModel.Delete()
	if err != nil {
		return 0, e.ErrInfo{
			Code: e.ErrorAiTokenDeleteFail,
			Msg:  err,
		}
	}
	return rows, e.ErrInfo{Code: e.SUCCESS}
}
