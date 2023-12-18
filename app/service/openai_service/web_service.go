package openai_service

import (
	"chatgpt_x/app/models/conversation"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// WebService OPENAI WEB 接口服务。
type WebService struct{}

// Save 保存对话记录，如果不存在则创建，存在则更新
func (s *WebService) Save(paramsModel conversation.Conversation) error {
	conversationModel, err := conversation.Get(paramsModel.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	// 对话信息不存在则创建，存在则更新
	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err = paramsModel.Create(); err != nil {
			return err
		}
		return nil
	}
	conversationModel.UpdateTime = time.Now()
	if _, err = conversationModel.Update(); err != nil {
		return err
	}
	return nil
}

// Conversation 平台对话。
func (s *WebService) Conversation(userID uint, body any) (<-chan []byte, error) {
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
	headers, err := GetBasicHeaders(aiTokenModel.Token, true)
	if err != nil {
		return nil, err
	}
	// 发送请求
	ch, err := SendStreamRequest("web", "POST", url, headers, body)
	if err != nil {
		return nil, err
	}
	// 存储对话记录
	conversationID, hookedCh := s.hookConversationID(ch)
	err = s.Save(conversation.Conversation{
		ID:        conversationID,
		UserID:    userID,
		AiTokenID: &aiTokenModel.ID,
		Type:      conversation.TypeWeb,
		ModelName: aiModelName.(string),
		Title:     "New chat",
		Status:    conversation.StatusEnable,
	})
	if err != nil {
		return nil, err
	}
	return hookedCh, nil
}

// hookConversationInfo 从数据包中提取 conversation_id。
func (s *WebService) hookConversationID(ch <-chan []byte) (string, <-chan []byte) {
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

// GetConversationHistory 获取对话历史。
func (s *WebService) GetConversationHistory(userID uint, offset, limit int) (string, error) {
	url := "/backend-api/conversations?offset=%s&limit=%s&order=updated"
	url = fmt.Sprintf(url, strconv.Itoa(offset), strconv.Itoa(limit))
	aiTokenModel, err := GetAiTokenFromUser(userID)
	if err != nil {
		return "", err
	}
	// 获取 headers
	header, err := GetBasicHeaders(aiTokenModel.Token, false)
	if err != nil {
		return "", err
	}
	// 发送请求
	result, err := SendRequest("web", "GET", url, header, nil)
	if err != nil {
		return "", err
	}
	return result, nil
}

// ChangeConversationTitle 修改对话标题。
func (s *WebService) ChangeConversationTitle(userID uint, conversationID string, body any) (string, error) {
	url := "/backend-api/conversation/" + conversationID
	aiTokenModel, err := GetAiTokenFromUser(userID)
	if err != nil {
		return "", err
	}
	// 获取当前用户的 token
	headers, err := GetBasicHeaders(aiTokenModel.Token, false)
	if err != nil {
		return "", err
	}
	// 发送请求
	result, err := SendRequest("web", "PATCH", url, headers, body)
	if err != nil {
		return "", err
	}
	return result, nil
}
