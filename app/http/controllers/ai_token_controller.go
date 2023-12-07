package controllers

import (
	"chatgpt_x/app/models/ai_token"
	"chatgpt_x/app/requests"
	"chatgpt_x/pkg/app"
	"chatgpt_x/pkg/e"
	"github.com/gin-gonic/gin"
	paginator "github.com/yafeng-Soong/gorm-paginator"
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
	// 检查密钥是否存在
	if ai_token.HasTokenExist(params.Token, 0) {
		appG.Response(http.StatusOK, e.ErrorAiTokenIsExist, nil, nil)
		return
	}
	// 创建密钥
	aiTokenModel := ai_token.AiToken{
		Type:   params.Type,
		Token:  params.Token,
		Remark: params.Remark,
		Status: params.Status,
	}
	if err := aiTokenModel.Create(); err != nil {
		appG.Response(http.StatusOK, e.ErrorAiTokenCreateFail, err, nil)
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
	// 检查密钥是否存在
	if ai_token.HasTokenExist(params.Token, int(params.ID)) {
		appG.Response(http.StatusOK, e.ErrorAiTokenIsExist, nil, nil)
		return
	}
	// 更新密钥
	aiTokenModel := ai_token.AiToken{
		ID:     params.ID,
		Type:   params.Type,
		Token:  params.Token,
		Remark: params.Remark,
		Status: params.Status,
	}
	rows, err := aiTokenModel.Update()
	if err != nil {
		appG.Response(http.StatusOK, e.ErrorAiTokenUpdateFail, err, nil)
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
	// 查询密钥列表
	aiTokenModel := ai_token.AiToken{}
	pageData, err := aiTokenModel.List(params.Page, params.PageSize)
	if err != nil {
		appG.Response(http.StatusOK, e.ErrorAiTokenSelectListFail, err, nil)
		return
	}
	data := pageData.(paginator.Page[ai_token.AiToken])
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
	// 删除密钥
	aiTokenModel := ai_token.AiToken{
		ID: params.ID,
	}
	rows, err := aiTokenModel.Delete()
	if err != nil {
		appG.Response(http.StatusOK, e.ErrorAiTokenDeleteFail, err, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, gin.H{"rows": rows})
}
