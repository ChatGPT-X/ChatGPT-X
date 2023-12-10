package controllers

import (
	"chatgpt_x/app/models/ai_token"
	"chatgpt_x/app/requests"
	"chatgpt_x/app/service/ai_token_service"
	"chatgpt_x/pkg/app"
	"chatgpt_x/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AiTokenController AI 密钥控制器。
type AiTokenController struct {
	BaseController
}

// Create 创建 AI 密钥。
func (at *AiTokenController) Create(c *gin.Context) {
	appG := at.GetAppG(c)
	// 表单验证
	var params requests.ValidateAiTokenCreate
	if err := c.ShouldBind(&params); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 创建 AI 密钥
	aiTokenService := ai_token_service.AiTokenService{}
	errInfo := aiTokenService.Create(ai_token.AiToken{
		Type:   params.Type,
		Token:  params.Token,
		Remark: params.Remark,
		Status: params.Status,
	})
	if errInfo.Code != e.SUCCESS {
		appG.Response(http.StatusOK, errInfo.Code, errInfo.Msg, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, nil)
}

// Update 更新 AI 密钥。
func (at *AiTokenController) Update(c *gin.Context) {
	appG := at.GetAppG(c)
	// 表单验证
	var params requests.ValidateAiTokenUpdate
	if err := c.ShouldBind(&params); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 更新 AI 密钥
	aiTokenService := ai_token_service.AiTokenService{}
	rows, errInfo := aiTokenService.Update(ai_token.AiToken{
		ID:     params.ID,
		Type:   params.Type,
		Token:  params.Token,
		Remark: params.Remark,
		Status: params.Status,
	})
	if errInfo.Code != e.SUCCESS {
		appG.Response(http.StatusOK, errInfo.Code, errInfo.Msg, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, gin.H{"rows": rows})
}

// List 查询 AI 密钥列表。
func (at *AiTokenController) List(c *gin.Context) {
	appG := at.GetAppG(c)
	// 表单验证
	var params requests.ValidateAiTokenList
	if err := c.ShouldBind(&params); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 设置默认值
	SetDefaultValue(&params.Page, 1)
	SetDefaultValue(&params.PageSize, 20)
	// 查询 AI 密钥列表
	aiTokenService := ai_token_service.AiTokenService{}
	data, errInfo := aiTokenService.List(params.Page, params.PageSize)
	if errInfo.Code != e.SUCCESS {
		appG.Response(http.StatusOK, errInfo.Code, errInfo.Msg, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, app.ResponseDataList{
		List:      data.Data,
		Page:      data.CurrentPage,
		PageSize:  data.PageSize,
		PageCount: data.Pages,
		Count:     data.Total,
	})
}

// Delete 删除 AI 密钥。
func (at *AiTokenController) Delete(c *gin.Context) {
	appG := at.GetAppG(c)
	// 表单验证
	var params requests.ValidateAiTokenDelete
	if err := c.ShouldBind(&params); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 删除 AI 密钥
	aiTokenService := ai_token_service.AiTokenService{}
	rows, errInfo := aiTokenService.Delete(params.ID)
	if errInfo.Code != e.SUCCESS {
		appG.Response(http.StatusOK, errInfo.Code, errInfo.Msg, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, gin.H{"rows": rows})
}
