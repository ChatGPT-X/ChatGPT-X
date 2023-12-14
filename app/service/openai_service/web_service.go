package openai_service

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// WebService OPENAI WEB 接口服务。
type WebService struct{}

// Conversation 平台对话。
func (s *WebService) Conversation(userID uint, body any) (<-chan []byte, error) {
	url := "/backend-api/conversation"
	// 获取 headers
	headers, err := GetBasicHeaders(userID, true)
	if err != nil {
		return nil, err
	}
	// 发送请求
	ch, err := SendStreamRequest("web", "POST", url, headers, body)
	if err != nil {
		return nil, err
	}
	conversationID, hookedCh := s.hookConversationID(ch)
	fmt.Println(conversationID)
	return hookedCh, nil
}

// hookConversationID 从数据包中提取 conversation_id。
func (s *WebService) hookConversationID(ch <-chan []byte) (string, <-chan []byte) {
	type conversationID struct {
		ConversationID string `json:"conversation_id"`
	}
	count := 3
	hookedCh := make(chan []byte, count)
	var jsonData conversationID
	// 同步处理前 count 个数据包，hook 拿 conversation_id
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
	// 获取 headers
	header, err := GetBasicHeaders(userID, false)
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
	// 获取当前用户的 token
	headers, err := GetBasicHeaders(userID, false)
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
