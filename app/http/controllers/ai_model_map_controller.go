package controllers

import (
	"chatgpt_x/app/models/ai_model_map"
	"chatgpt_x/app/requests"
	"chatgpt_x/pkg/app"
	"chatgpt_x/pkg/e"
	"github.com/gin-gonic/gin"
	paginator "github.com/yafeng-Soong/gorm-paginator"
	"net/http"
)

// AiModelMapController AI 模型关系映射控制器。
type AiModelMapController struct {
	BaseController
}

// Create 创建 AI 模型关系映射。
func (am *AiModelMapController) Create(c *gin.Context) {
	appG := am.GetAppG(c)
	// 表单验证
	var form requests.ValidateAiModelMapCreate
	if err := c.ShouldBind(&form); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 检查 AI 模型是否存在
	if ai_model_map.HasAiModelExist(form.AiName, 0) {
		appG.Response(http.StatusOK, e.ErrorAiModelIsExist, nil, nil)
		return
	}
	// 创建 AI 模型关系映射
	aiModelMap := ai_model_map.AiModelMap{
		DisplayName: form.DisplayName,
		AiName:      form.AiName,
		IsDisabled:  form.IsDisabled,
	}
	if err := aiModelMap.Create(); err != nil {
		appG.Response(http.StatusOK, e.ErrorAiModelMapCreateFail, err, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, nil)
}

// Update 更新 AI 模型关系映射。
func (am *AiModelMapController) Update(c *gin.Context) {
	appG := am.GetAppG(c)
	// 表单验证
	var form requests.ValidateAiModelMapUpdate
	if err := c.ShouldBind(&form); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 检查 AI 模型是否存在
	if ai_model_map.HasAiModelExist(form.AiName, int(form.ID)) {
		appG.Response(http.StatusOK, e.ErrorAiModelIsExist, nil, nil)
		return
	}
	// 更新 AI 模型关系映射
	aiModelMap := ai_model_map.AiModelMap{
		ID:          form.ID,
		DisplayName: form.DisplayName,
		AiName:      form.AiName,
		IsDisabled:  form.IsDisabled,
	}
	rows, err := aiModelMap.Update()
	if err != nil {
		appG.Response(http.StatusOK, e.ErrorAiModelMapUpdateFail, err, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, gin.H{"rows": rows})
}

// Select 查询 AI 模型关系映射。
func (am *AiModelMapController) Select(c *gin.Context) {
	appG := am.GetAppG(c)
	// 表单验证
	var form requests.ValidateAiModelMapSelect
	if err := c.ShouldBind(&form); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 设置默认值
	SetDefaultValue(&form.Page, 1)
	SetDefaultValue(&form.PageSize, 20)
	// 查询 AI 模型关系映射
	aiModelMap := ai_model_map.AiModelMap{}
	pageData, err := aiModelMap.Select(form.Page, form.PageSize)
	if err != nil {
		appG.Response(http.StatusOK, e.ErrorAiModelMapSelectFail, err, nil)
		return
	}
	data := pageData.(paginator.Page[ai_model_map.AiModelMap])
	appG.Response(http.StatusOK, e.SUCCESS, nil, app.ResponseDataList{
		List:      data.Data,
		Page:      data.CurrentPage,
		PageSize:  data.PageSize,
		PageCount: data.Pages,
		Count:     data.Total,
	})
}

// Delete 删除 AI 模型关系映射。
func (am *AiModelMapController) Delete(c *gin.Context) {
	appG := am.GetAppG(c)
	// 表单验证
	var form requests.ValidateAiModelMapDelete
	if err := c.ShouldBind(&form); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 删除 AI 模型关系映射
	aiModelMap := ai_model_map.AiModelMap{
		ID: form.ID,
	}
	rows, err := aiModelMap.Delete()
	if err != nil {
		appG.Response(http.StatusOK, e.ErrorAiModelMapDeleteFail, err, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, gin.H{"rows": rows})
}
