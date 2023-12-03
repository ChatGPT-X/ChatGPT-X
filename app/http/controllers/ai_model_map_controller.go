package controllers

import (
	"chatgpt_x/app/models/ai_model_map"
	"chatgpt_x/app/requests"
	"chatgpt_x/pkg/e"
	"github.com/gin-gonic/gin"
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
