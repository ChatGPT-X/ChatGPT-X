package controllers

import (
	"chatgpt_x/app/models/token"
	"chatgpt_x/app/requests"
	"chatgpt_x/pkg/e"
	"github.com/gin-gonic/gin"
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
