package controllers

import (
	"chatgpt_x/app/models/user"
	"chatgpt_x/app/requests"
	"chatgpt_x/pkg/auth"
	"chatgpt_x/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// UserController 用户控制器。
type UserController struct {
	BaseController
}

// DoRegister 用户注册。
func (u *UserController) DoRegister(c *gin.Context) {
	appG := u.GetAppG(c)
	// 表单验证
	var form requests.ValidateDoRegister
	if err := c.ShouldBind(&form); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 检查用户名是否重复
	if user.HasByUsernameExist(form.Username) {
		appG.Response(http.StatusOK, e.ErrorUserIsExist, nil, nil)
		return
	}
	// 创建用户
	userModel := user.Users{
		Username:      form.Username,
		Email:         form.Email,
		Password:      form.Password,
		LastLoginTime: time.Now(),
	}
	if err := userModel.Create(); err != nil {
		appG.Response(http.StatusOK, e.ErrorUserCreateFail, err, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, nil)
}

// DoLogin 用户登录。
func (u *UserController) DoLogin(c *gin.Context) {
	appG := u.GetAppG(c)
	// 表单验证
	var form requests.ValidateDoLogin
	if err := c.ShouldBind(&form); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 检查用户名和密码是否正确
	userModel, err := user.GetByUsername(form.Username)
	if err != nil || !user.CheckPassword(form.Password, userModel.Password) {
		appG.Response(http.StatusOK, e.ErrorIncorrectUsernameOrPassword, err, nil)
		return
	}
	// 检查用户是否被封禁
	if userModel.IsDisable() {
		appG.Response(http.StatusOK, e.ErrorUserIsDisabled, nil, nil)
		return
	}
	// 生成 Token 授权
	jwt, err := auth.GenerateToken(auth.CustomClaims{
		UserID:   userModel.ID,
		IsAdmin:  userModel.IsAdmin,
		Username: userModel.Username,
		Email:    userModel.Email,
	})
	if err != nil {
		appG.Response(http.StatusOK, e.ErrorGenerateTokenFail, err, nil)
		return
	}
	c.Header("Authorization", "Bearer "+jwt)
	appG.Response(http.StatusOK, e.SUCCESS, nil, nil)
}

// Logout 用户登出。
func (u *UserController) Logout(c *gin.Context) {
	appG := u.GetAppG(c)
	c.Header("Authorization", "")
	appG.Response(http.StatusOK, e.SUCCESS, nil, nil)
}
