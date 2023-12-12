package controllers

import (
	"chatgpt_x/app/service/openai_service"
	"chatgpt_x/pkg/e"
	"chatgpt_x/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// OpenaiWebController OPENAI WEB 接口控制器。
type OpenaiWebController struct {
	BaseController
}

// Conversation WEB 平台对话。
func (ow *OpenaiWebController) Conversation(c *gin.Context) {
	appG := ow.GetAppG(c)
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	// 接收参数
	var params any
	err := c.ShouldBindJSON(&params)
	if err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	userID := getUserID(c)
	openaiService := openai_service.WebService{}
	ch, err := openaiService.Conversation(userID, params)
	if err != nil {
		logger.Error(err)
	}
	for b := range ch {
		fmt.Fprintf(c.Writer, "%s\n\n", string(b))
	}
}
