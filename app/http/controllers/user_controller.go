package controllers

import (
	"chatgpt_x/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AdminController 书籍控制器.
type AdminController struct {
	BaseController
}

func (a *AdminController) DoLogin(c *gin.Context) {
	appG := a.GetAppG(c)
	//username := c.PostForm("username")
	//password := c.PostForm("password")

	appG.Response(http.StatusOK, e.SUCCESS, nil, nil)
}
