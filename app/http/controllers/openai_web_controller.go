package controllers

import (
	"chatgpt_x/app/requests"
	"chatgpt_x/app/service/openai_service"
	"chatgpt_x/pkg/e"
	"chatgpt_x/pkg/logger"
	"encoding/json"
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
		c.SSEvent("", string(b))
	}
}

// GetConversationHistory 获取对话历史。
func (ow *OpenaiWebController) GetConversationHistory(c *gin.Context) {
	appG := ow.GetAppG(c)
	// 接收参数
	var params requests.ValidateGetConversationHistory
	if err := c.ShouldBindQuery(&params); err != nil {
		appG.Response(http.StatusOK, e.InvalidParams, err, nil)
		return
	}
	SetDefaultValue(&params.Offset, 0)
	SetDefaultValue(&params.Limit, 28)
	userID := getUserID(c)
	openaiService := openai_service.WebService{}
	result, err := openaiService.GetConversationHistory(userID, params.Offset, params.Limit)
	if err != nil {
		appG.Response(http.StatusOK, e.ErrorGetConversationHistoryFail, err, nil)
		return
	}
	var data any
	json.Unmarshal([]byte(result), &data)
	appG.Response(http.StatusOK, e.SUCCESS, nil, data)
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
