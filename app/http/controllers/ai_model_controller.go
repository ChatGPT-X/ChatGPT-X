package controllers

import (
	"chatgpt_x/app/models/ai_model"
	"chatgpt_x/app/requests"
	"chatgpt_x/pkg/app"
	"chatgpt_x/pkg/e"
	"github.com/gin-gonic/gin"
	paginator "github.com/yafeng-Soong/gorm-paginator"
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
	var form requests.ValidateAiModelCreate
	if err := c.ShouldBind(&form); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 检查 AI 模型是否存在
	if ai_model.HasAiModelExist(form.AiName, 0) {
		appG.Response(http.StatusOK, e.ErrorAiModelIsExist, nil, nil)
		return
	}
	// 创建 AI 模型
	aiModel := ai_model.AiModel{
		DisplayName: form.DisplayName,
		AiName:      form.AiName,
		Status:      form.Status,
	}
	if err := aiModel.Create(); err != nil {
		appG.Response(http.StatusOK, e.ErrorAiModelCreateFail, err, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, nil)
}

// Update 更新 AI 模型。
func (am *AiModelController) Update(c *gin.Context) {
	appG := am.GetAppG(c)
	// 表单验证
	var form requests.ValidateAiModelUpdate
	if err := c.ShouldBind(&form); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 检查 AI 模型是否存在
	if ai_model.HasAiModelExist(form.AiName, int(form.ID)) {
		appG.Response(http.StatusOK, e.ErrorAiModelIsExist, nil, nil)
		return
	}
	// 更新 AI 模型
	aiModel := ai_model.AiModel{
		ID:          form.ID,
		DisplayName: form.DisplayName,
		AiName:      form.AiName,
		Status:      form.Status,
	}
	rows, err := aiModel.Update()
	if err != nil {
		appG.Response(http.StatusOK, e.ErrorAiModelUpdateFail, err, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, gin.H{"rows": rows})
}

// List 查询 AI 模型列表。
func (am *AiModelController) List(c *gin.Context) {
	appG := am.GetAppG(c)
	// 表单验证
	var form requests.ValidateAiModelList
	if err := c.ShouldBind(&form); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 设置默认值
	SetDefaultValue(&form.Page, 1)
	SetDefaultValue(&form.PageSize, 20)
	// 查询 AI 模型列表
	aiModel := ai_model.AiModel{}
	pageData, err := aiModel.List(form.Page, form.PageSize)
	if err != nil {
		appG.Response(http.StatusOK, e.ErrorAiModelSelectFail, err, nil)
		return
	}
	data := pageData.(paginator.Page[ai_model.AiModel])
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
	var form requests.ValidateAiModelDelete
	if err := c.ShouldBind(&form); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 删除 AI 模型
	aiModel := ai_model.AiModel{
		ID: form.ID,
	}
	rows, err := aiModel.Delete()
	if err != nil {
		appG.Response(http.StatusOK, e.ErrorAiModelDeleteFail, err, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, gin.H{"rows": rows})
}
