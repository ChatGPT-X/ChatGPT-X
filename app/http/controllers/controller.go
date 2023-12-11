package controllers

import (
	"chatgpt_x/pkg/app"
	"chatgpt_x/pkg/auth"
	"github.com/gin-gonic/gin"
	"reflect"
)

// BaseController 基础组件控制器。
type BaseController struct{}

// GetAppG 获取一个 appGin 实例。
func (b *BaseController) GetAppG(c *gin.Context) *app.Gin {
	return &app.Gin{C: c}
}

// SetDefaultValue 设置默认值函数。
func SetDefaultValue[T comparable](param *T, defaultValue T) {
	if *param == reflect.Zero(reflect.TypeOf(*param)).Interface() {
		*param = defaultValue
	}
}

// getUserID 获取用户 ID。
func getUserID(c *gin.Context) (userID uint) {
	jwt, _ := auth.GetTokenFromHeader(c)
	claims, _ := auth.ParseToken(jwt)
	return claims.UserID
}
