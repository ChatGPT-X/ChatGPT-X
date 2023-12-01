package controllers

import (
	"chatgpt_x/app/models/user"
	"chatgpt_x/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserController 用户控制器。
type UserController struct {
	BaseController
}

// DoRegisterForm 注册表单。
type doRegisterForm struct {
	Username string `form:"username" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// DoRegister 用户注册。
func (u *UserController) DoRegister(c *gin.Context) {
	appG := u.GetAppG(c)
	// 表单验证
	var form doRegisterForm
	if err := c.ShouldBind(&form); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 检查用户名是否重复
	if user.HasByUsername(form.Username) {
		appG.Response(http.StatusOK, e.ErrorUserIsExist, nil, nil)
		return
	}
	// 创建用户
	userModel := user.User{
		Username: form.Username,
		Email:    form.Email,
		Password: form.Password,
	}
	if err := userModel.Create(); err != nil {
		appG.Response(http.StatusOK, e.ErrorUserCreateFail, err, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, nil)
}

type doLoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// DoLogin 用户登录。
func (u *UserController) DoLogin(c *gin.Context) {
	appG := u.GetAppG(c)
	session := u.GetSessions(c)
	// 表单验证
	var form doLoginForm
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
	if userModel.IsDisabled() {
		appG.Response(http.StatusOK, e.ErrorUserIsDisabled, nil, nil)
		return
	}
	// 保存用户信息到 Session
	info := map[string]interface{}{
		"user_id":  userModel.ID,
		"email":    userModel.Email,
		"username": userModel.Username,
	}
	session.Set("user_info", info)
	_ = session.Save()
	appG.Response(http.StatusOK, e.SUCCESS, nil, nil)
}
