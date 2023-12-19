package controllers

import (
	"chatgpt_x/app/requests"
	"chatgpt_x/app/service/openai_service"
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
	// 接收参数
	var paramsJson any
	err := c.ShouldBindJSON(&paramsJson)
	if err != nil {
		appG.ResponseWithOpenai(http.StatusInternalServerError, "Params error, please contact the developer.")
		return
	}
	userID := getUserID(c)
	openaiService := openai_service.WebService{}
	respResult, err := openaiService.Conversation(userID, paramsJson)
	if err != nil {
		appG.ResponseWithOpenai(http.StatusInternalServerError, "OpenAI error, please wait and try again.")
		return
	}
	c.Header("Connection", "keep-alive")
	c.Header("Cache-Control", "no-cache")
	c.Header("Content-Type", respResult.BodyType)
	switch respResult.BodyType {
	// 返回 json 格式则代表出错误了
	case "application/json", "application/json; charset=utf-8":
		var result any
		_ = json.Unmarshal(respResult.Body, &result)
		appG.ResponseWithOpenai(respResult.StatusCode, string(respResult.Body))
		return
	// 正常是 event-stream 流式传输
	case "text/event-stream", "text/event-stream; charset=utf-8":
		for b := range respResult.BodyStream {
			c.SSEvent("", string(b))
		}
	}
}

// GetConversationHistory 获取对话历史。
func (ow *OpenaiWebController) GetConversationHistory(c *gin.Context) {
	appG := ow.GetAppG(c)
	// 接收参数
	var params requests.ValidateGetConversationHistory
	if err := c.ShouldBindQuery(&params); err != nil {
		appG.ResponseWithOpenai(http.StatusInternalServerError, "Params error, please contact the developer.")
		return
	}
	SetDefaultValue(&params.Offset, 0)
	SetDefaultValue(&params.Limit, 28)
	userID := getUserID(c)
	openaiService := openai_service.WebService{}
	respResult, err := openaiService.GetConversationHistory(userID, params.Offset, params.Limit)
	if err != nil {
		appG.ResponseWithOpenai(http.StatusInternalServerError, "OpenAI error, please wait and try again.")
		return
	}
	var result any
	_ = json.Unmarshal(respResult.Body, &result)
	c.JSON(respResult.StatusCode, result)
}

// ChangeConversationTitle 修改对话标题。
func (ow *OpenaiWebController) ChangeConversationTitle(c *gin.Context) {
	appG := ow.GetAppG(c)
	// 接收参数
	var params requests.ValidateUUIDv4
	if err := c.ShouldBindUri(&params); err != nil {
		appG.ResponseWithOpenai(http.StatusInternalServerError, "Params error, please contact the developer.")
		return
	}
	var paramsJson any
	err := c.ShouldBindJSON(&paramsJson)
	if err != nil {
		appG.ResponseWithOpenai(http.StatusInternalServerError, "Params error, please contact the developer.")
		return
	}
	userID := getUserID(c)
	openaiService := openai_service.WebService{}
	respResult, err := openaiService.ChangeConversationTitle(userID, params.ConversationID, paramsJson)
	if err != nil {
		appG.ResponseWithOpenai(http.StatusInternalServerError, "OpenAI error, please wait and try again.")
		return
	}
	var result any
	_ = json.Unmarshal(respResult.Body, &result)
	c.JSON(respResult.StatusCode, result)
}
