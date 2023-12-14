package controllers

import (
	"chatgpt_x/app/models"
	"chatgpt_x/app/models/user"
	"chatgpt_x/app/requests"
	"chatgpt_x/app/service/user_service"
	"chatgpt_x/pkg/app"
	"chatgpt_x/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserController 用户控制器。
type UserController struct {
	BaseController
}

// DoRegister 用户注册。
func (u *UserController) DoRegister(c *gin.Context) {
	appG := u.GetAppG(c)
	// 表单验证
	var params requests.ValidateDoRegister
	if err := c.ShouldBind(&params); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 用户注册
	userService := user_service.UserService{}
	errInfo := userService.DoRegister(user.User{
		AiTokenID: models.SqlNullUint,
		Username:  params.Username,
		Email:     params.Email,
		Password:  params.Password,
	})
	if errInfo.Code != e.SUCCESS {
		appG.Response(http.StatusOK, errInfo.Code, errInfo.Msg, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, nil)
}

// DoLogin 用户登录。
func (u *UserController) DoLogin(c *gin.Context) {
	appG := u.GetAppG(c)
	// 表单验证
	var params requests.ValidateDoLogin
	if err := c.ShouldBind(&params); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 用户登录
	userService := user_service.UserService{}
	token, errInfo := userService.DoLogin(user.User{
		Username: params.Username,
		Password: params.Password,
	})
	if errInfo.Code != e.SUCCESS {
		appG.Response(http.StatusOK, errInfo.Code, errInfo.Msg, nil)
		return
	}
	appG.SetAuthorization(token)
	appG.Response(http.StatusOK, e.SUCCESS, nil, nil)
}

// Logout 用户登出。
func (u *UserController) Logout(c *gin.Context) {
	appG := u.GetAppG(c)
	appG.SetAuthorization("")
	appG.Response(http.StatusOK, e.SUCCESS, nil, nil)
}

// List 获取用户列表。
func (u *UserController) List(c *gin.Context) {
	appG := u.GetAppG(c)
	// 表单验证
	var params requests.ValidateUserList
	if err := c.ShouldBind(&params); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 设置默认值
	SetDefaultValue(&params.Page, 1)
	SetDefaultValue(&params.PageSize, 20)
	// 获取用户列表
	userService := user_service.UserService{}
	data, errInfo := userService.List(params.Page, params.PageSize)
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

// Update 更新用户。
func (u *UserController) Update(c *gin.Context) {
	appG := u.GetAppG(c)
	// 表单验证
	var params requests.ValidateUserUpdate
	if err := c.ShouldBind(&params); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 更新用户
	userService := user_service.UserService{}
	rows, errInfo := userService.Update(user.User{
		ID:        params.ID,
		AiTokenID: params.AiTokenID,
		Password:  params.Password,
		Status:    params.Status,
	})
	if errInfo.Code != e.SUCCESS {
		appG.Response(http.StatusOK, errInfo.Code, errInfo.Msg, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, gin.H{"rows": rows})
}

// Delete 删除用户。
func (u *UserController) Delete(c *gin.Context) {
	appG := u.GetAppG(c)
	// 表单验证
	var params requests.ValidateUserDelete
	if err := c.ShouldBind(&params); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 删除用户
	userService := user_service.UserService{}
	rows, errInfo := userService.Delete(params.ID)
	if errInfo.Code != e.SUCCESS {
		appG.Response(http.StatusOK, errInfo.Code, errInfo.Msg, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, gin.H{"rows": rows})
}
