package openai_service

import (
	"chatgpt_x/app/models/conversation"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// WebService OPENAI WEB 接口服务。
type WebService struct{}

// Save 保存对话记录，如果不存在则创建，存在则更新
func (s *WebService) Save(conversationModel conversation.Conversation) error {
	// 检查是否已经存在
	return conversationModel.CreateOrUpdate()
}

// Conversation 平台对话。
func (s *WebService) Conversation(userID uint, body any) (*ResponseResult, error) {
	url := "/backend-api/conversation"
	aiModelName, ok := body.(map[string]any)["model"]
	if !ok {
		return nil, fmt.Errorf("model is not exist")
	}
	aiTokenModel, err := GetAiTokenFromUser(userID)
	if err != nil {
		return nil, err
	}
	// 获取 headers
	headers := GetBasicHeaders(aiTokenModel.Token, true)
	// 发送请求
	respResult, err := SendStreamRequest("web", "POST", url, headers, body)
	if err != nil {
		return respResult, err
	}
	switch respResult.BodyType {
	case "text/event-stream", "text/event-stream; charset=utf-8":
		// 存储对话记录
		conversationID, hookedCh := s.hookConversationID(respResult.BodyStream)
		respResult.BodyStream = hookedCh
		err = s.Save(conversation.Conversation{
			UserID:            userID,
			AiTokenID:         &aiTokenModel.ID,
			Type:              conversation.TypeWeb,
			ModelName:         aiModelName.(string),
			ConversationID:    conversationID,
			ConversationTitle: "New chat",
			Status:            conversation.StatusEnable,
			UpdateTime:        time.Now(),
		})
		if err != nil {
			return nil, err
		}
	}
	return respResult, nil
}

// hookConversationID 从数据包中提取 conversation_id。
func (s *WebService) hookConversationID(ch <-chan []byte) (string, chan []byte) {
	type conversationID struct {
		ConversationID string `json:"conversation_id"`
	}
	count := 3
	hookedCh := make(chan []byte, count)
	var jsonData conversationID
	// 同步处理前 count 个数据包拿到 conversation_id
	for count > 0 {
		data, ok := <-ch
		if !ok {
			break
		}
		err := json.Unmarshal(data, &jsonData)
		if err != nil || (err == nil && jsonData.ConversationID == "") {
			hookedCh <- data
			count--
			continue
		}
		hookedCh <- data
		break
	}
	// 异步处理剩余的数据包
	go func() {
		defer close(hookedCh)
		for data := range ch {
			hookedCh <- data
		}
	}()
	return jsonData.ConversationID, hookedCh
}

// ConversationHistory 对话历史接口要返回的结构。
type ConversationHistory struct {
	HasMissingConversations bool                        `json:"has_missing_conversations"`
	Items                   []conversation.Conversation `json:"items"`
	Limit                   int                         `json:"limit"`
	Offset                  int                         `json:"offset"`
	Total                   int64                       `json:"total"`
}

// GetConversationHistory 获取对话历史。
func (s *WebService) GetConversationHistory(userID uint, offset, limit int) (ConversationHistory, error) {
	history, count, err := conversation.List(offset, limit, userID, conversation.StatusEnable)
	if err != nil {
		return ConversationHistory{}, err
	}
	conversationHistory := ConversationHistory{
		HasMissingConversations: false,
		Items:                   history,
		Limit:                   limit,
		Offset:                  offset,
		Total:                   count,
	}
	return conversationHistory, nil
}

// ChangeConversationTitle 修改对话标题。
func (s *WebService) ChangeConversationTitle(userID uint, conversationID string, body any) (*ResponseResult, error) {
	url := "/backend-api/conversation/" + conversationID
	aiTokenModel, err := GetAiTokenFromUser(userID)
	if err != nil {
		return nil, err
	}
	// 获取当前用户的 token
	headers := GetBasicHeaders(aiTokenModel.Token, false)
	// 发送请求
	respResult, err := SendRequest("web", "PATCH", url, headers, body)
	if err != nil {
		return nil, err
	}
	// 如果请求成功则更新对话标题
	if respResult.StatusCode == http.StatusOK {
		title, ok := body.(map[string]any)["title"]
		if !ok {
			return nil, fmt.Errorf("title is not exist")
		}
		conversationModel := conversation.Conversation{
			ConversationID:    conversationID,
			ConversationTitle: title.(string),
		}
		if _, err = conversationModel.UpdateByConversationID(); err != nil {
			return nil, err
		}
	}
	return respResult, nil
}
