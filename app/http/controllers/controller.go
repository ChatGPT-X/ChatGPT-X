package controllers

import (
	"chatgpt_x/pkg/app"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// BaseController 基础组件控制器.
type BaseController struct{}

// GetAppG 获取一个 appGin 实例.
func (b *BaseController) GetAppG(c *gin.Context) *app.Gin {
	return &app.Gin{C: c}
}

// GetSessions 获取一个 Sessions 实例.
func (b *BaseController) GetSessions(c *gin.Context) sessions.Session {
	return sessions.Default(c)
}
