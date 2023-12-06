package controllers

import (
	"chatgpt_x/app/models/token"
	"chatgpt_x/app/requests"
	"chatgpt_x/pkg/app"
	"chatgpt_x/pkg/e"
	"github.com/gin-gonic/gin"
	paginator "github.com/yafeng-Soong/gorm-paginator"
	"net/http"
)

// TokenController 密钥控制器。
type TokenController struct {
	BaseController
}

// Create 创建密钥。
func (t *TokenController) Create(c *gin.Context) {
	appG := t.GetAppG(c)
	// 表单验证
	var params requests.ValidateTokenCreate
	if err := c.ShouldBind(&params); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 检查密钥是否存在
	if token.HasTokenExist(params.Token, 0) {
		appG.Response(http.StatusOK, e.ErrorTokenIsExist, nil, nil)
		return
	}
	// 创建密钥
	tokenModel := token.Token{
		Type:   params.Type,
		Token:  params.Token,
		Remark: params.Remark,
		Status: params.Status,
	}
	if err := tokenModel.Create(); err != nil {
		appG.Response(http.StatusOK, e.ErrorTokenCreateFail, err, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, nil)
}

// Update 更新密钥。
func (t *TokenController) Update(c *gin.Context) {
	appG := t.GetAppG(c)
	// 表单验证
	var params requests.ValidateTokenUpdate
	if err := c.ShouldBind(&params); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 检查密钥是否存在
	if token.HasTokenExist(params.Token, int(params.ID)) {
		appG.Response(http.StatusOK, e.ErrorTokenIsExist, nil, nil)
		return
	}
	// 更新密钥
	tokenModel := token.Token{
		ID:     params.ID,
		Type:   params.Type,
		Token:  params.Token,
		Remark: params.Remark,
		Status: params.Status,
	}
	rows, err := tokenModel.Update()
	if err != nil {
		appG.Response(http.StatusOK, e.ErrorTokenUpdateFail, err, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, gin.H{"rows": rows})
}

// List 查询密钥列表。
func (t *TokenController) List(c *gin.Context) {
	appG := t.GetAppG(c)
	// 表单验证
	var params requests.ValidateTokenList
	if err := c.ShouldBind(&params); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 设置默认值
	SetDefaultValue(&params.Page, 1)
	SetDefaultValue(&params.PageSize, 20)
	// 查询 AI 模型列表
	tokenModel := token.Token{}
	pageData, err := tokenModel.List(params.Page, params.PageSize)
	if err != nil {
		appG.Response(http.StatusOK, e.ErrorTokenSelectListFail, err, nil)
		return
	}
	data := pageData.(paginator.Page[token.Token])
	appG.Response(http.StatusOK, e.SUCCESS, nil, app.ResponseDataList{
		List:      data.Data,
		Page:      data.CurrentPage,
		PageSize:  data.PageSize,
		PageCount: data.Pages,
		Count:     data.Total,
	})
}
