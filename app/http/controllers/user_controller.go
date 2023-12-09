package controllers

import (
	"chatgpt_x/app/models"
	"chatgpt_x/app/models/user"
	"chatgpt_x/app/requests"
	"chatgpt_x/pkg/app"
	"chatgpt_x/pkg/auth"
	"chatgpt_x/pkg/e"
	"github.com/gin-gonic/gin"
	paginator "github.com/yafeng-Soong/gorm-paginator"
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
	var params requests.ValidateDoRegister
	if err := c.ShouldBind(&params); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 检查用户名是否重复
	if user.HasByUsernameExist(params.Username) {
		appG.Response(http.StatusOK, e.ErrorUserIsExist, nil, nil)
		return
	}
	// 创建用户
	userModel := user.User{
		AiTokenID: models.SqlNullUint,
		Username:  params.Username,
		Email:     params.Email,
		Password:  params.Password,
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
	var params requests.ValidateDoLogin
	if err := c.ShouldBind(&params); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	// 检查用户名和密码是否正确
	userModel, err := user.GetByUsername(params.Username)
	if err != nil || !user.CheckPassword(params.Password, userModel.Password) {
		appG.Response(http.StatusOK, e.ErrorIncorrectUsernameOrPassword, err, nil)
		return
	}
	// 检查用户是否被封禁
	if userModel.Status == user.StatusDisable {
		appG.Response(http.StatusOK, e.ErrorUserIsDisabled, nil, nil)
		return
	}
	// 更新用户登录时间
	userModel.LastLoginTime = time.Now()
	if _, err = userModel.Update(); err != nil {
		appG.Response(http.StatusOK, e.ErrorUserLoginFail, err, nil)
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

// List 查询用户列表。
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
	// 查询用户列表
	userModel := user.User{}
	pageData, err := userModel.List(params.Page, params.PageSize)
	if err != nil {
		appG.Response(http.StatusOK, e.ErrorUserSelectListFail, err, nil)
		return
	}
	data := pageData.(paginator.Page[user.User])
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
	userModel, err := user.Get(params.ID)
	if err != nil {
		appG.Response(http.StatusOK, e.ErrorUserSelectDetailFail, err, nil)
		return
	}
	userModel.AiTokenID = params.AiTokenID
	userModel.Password = params.Password
	userModel.Status = params.Status
	rows, err := userModel.Update()
	if err != nil {
		appG.Response(http.StatusOK, e.ErrorUserUpdateFail, err, nil)
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
	userModel := user.User{
		ID: params.ID,
	}
	rows, err := userModel.Delete()
	if err != nil {
		appG.Response(http.StatusOK, e.ErrorUserDeleteFail, err, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil, gin.H{"rows": rows})
}
