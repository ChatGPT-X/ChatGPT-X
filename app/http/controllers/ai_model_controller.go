package controllers

import (
	"chatgpt_x/app/models/ai_model"
	"chatgpt_x/app/requests"
	"chatgpt_x/app/service/ai_model_service"
	"chatgpt_x/pkg/app"
	"chatgpt_x/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AiModelController AI 模型控制器。
type AiModelController struct {
	BaseController
}

// Create 创建 AI 模型。
func (am *AiModelController) Create(c *gin.Context) {
	appG := am.GetAppG(c)
	// 表单验证
	var params requests.ValidateAiModelCreate
	if err := c.ShouldBind(&params); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 创建 AI 模型
	aiModelService := ai_model_service.AiModelService{}
	errInfo := aiModelService.Create(ai_model.AiModel{
		Type:      params.Type,
		AliasName: params.AliasName,
		Name:      params.Name,
		Status:    params.Status,
	})
	if errInfo.Code != e.SUCCESS {
		appG.Response(http.StatusOK, errInfo.Code, errInfo.Msg, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, nil)
}

// Update 更新 AI 模型。
func (am *AiModelController) Update(c *gin.Context) {
	appG := am.GetAppG(c)
	// 表单验证
	var params requests.ValidateAiModelUpdate
	if err := c.ShouldBind(&params); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 更新 AI 模型
	aiModelService := ai_model_service.AiModelService{}
	rows, errInfo := aiModelService.Update(ai_model.AiModel{
		ID:        params.ID,
		Type:      params.Type,
		AliasName: params.AliasName,
		Name:      params.Name,
		Status:    params.Status,
	})
	if errInfo.Code != e.SUCCESS {
		appG.Response(http.StatusOK, errInfo.Code, errInfo.Msg, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, gin.H{"rows": rows})
}

// List 获取 AI 模型列表。
func (am *AiModelController) List(c *gin.Context) {
	appG := am.GetAppG(c)
	// 表单验证
	var params requests.ValidateAiModelList
	if err := c.ShouldBind(&params); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 设置默认值
	SetDefaultValue(&params.Page, 1)
	SetDefaultValue(&params.PageSize, 20)
	// 获取 AI 模型列表
	aiModelService := ai_model_service.AiModelService{}
	data, errInfo := aiModelService.List(params.Type, params.Page, params.PageSize)
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

// Delete 删除 AI 模型。
func (am *AiModelController) Delete(c *gin.Context) {
	appG := am.GetAppG(c)
	// 表单验证
	var params requests.ValidateAiModelDelete
	if err := c.ShouldBind(&params); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 删除 AI 模型
	aiModelService := ai_model_service.AiModelService{}
	rows, errInfo := aiModelService.Delete(params.ID)
	if errInfo.Code != e.SUCCESS {
		appG.Response(http.StatusOK, errInfo.Code, errInfo.Msg, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, gin.H{"rows": rows})
}
