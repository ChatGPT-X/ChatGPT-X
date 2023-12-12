package controllers

import (
	"chatgpt_x/app/requests"
	"chatgpt_x/app/service/openai_service"
	"chatgpt_x/pkg/e"
	"chatgpt_x/pkg/logger"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// OpenaiWebController OPENAI WEB 接口控制器。
type OpenaiWebController struct {
	BaseController
}

// Conversation 平台对话。
func (ow *OpenaiWebController) Conversation(c *gin.Context) {
	appG := ow.GetAppG(c)
	c.Header("Connection", "keep-alive")
	c.Header("Cache-Control", "no-cache")
	c.Header("Content-Type", "text/event-stream")
	// 接收参数
	var paramsJson any
	err := c.ShouldBindJSON(&paramsJson)
	if err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	userID := getUserID(c)
	openaiService := openai_service.WebService{}
	ch, err := openaiService.Conversation(userID, paramsJson)
	if err != nil {
		logger.Error(err)
		return
	}
	for b := range ch {
		fmt.Fprintf(c.Writer, "%s\n\n", string(b))
	}
}

// ChangeConversationTitle 修改对话标题。
func (ow *OpenaiWebController) ChangeConversationTitle(c *gin.Context) {
	appG := ow.GetAppG(c)
	// 接收参数
	var params requests.ValidateUUIDv4
	if err := c.ShouldBindUri(&params); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	var paramsJson any
	err := c.ShouldBindJSON(&paramsJson)
	if err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	userID := getUserID(c)
	openaiService := openai_service.WebService{}
	result, err := openaiService.ChangeConversationTitle(userID, params.ConversationID, paramsJson)
	if err != nil {
		appG.Response(http.StatusOK, e.ErrorChangeConversationTitleFail, err, nil)
		return
	}
	var data any
	json.Unmarshal([]byte(result), &data)
	appG.Response(http.StatusOK, e.SUCCESS, nil, data)
}
