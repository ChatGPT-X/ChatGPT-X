package openai_service

// WebService OPENAI WEB 接口服务。
type WebService struct{}

// Conversation 平台对话。
func (s *WebService) Conversation(userID uint, body any) (<-chan []byte, error) {
	// 获取当前用户的 token
	aiTokenModel, err := GetAiTokenFromUser(userID)
	if err != nil {
		return nil, err
	}
	url := "/backend-api/conversation"
	headers := map[string]string{
		"Authorization": "Bearer " + aiTokenModel.Token,
		"Content-Type":  "application/json; charset=utf-8",
		"User-Agent":    "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Referer":       "https://chat.openai.com",
		"Origin":        "https://chat.openai.com",
		"Accept":        "text/event-stream",
		"Cache-Control": "no-cache",
	}
	ch, err := SendStreamRequest("web", "POST", url, headers, body)
	if err != nil {
		return nil, err
	}
	return ch, nil
}

// ChangeConversationTitle 修改对话标题。
func (s *WebService) ChangeConversationTitle(userID uint, conversationID string, body any) (string, error) {
	// 获取当前用户的 token
	aiTokenModel, err := GetAiTokenFromUser(userID)
	if err != nil {
		return "", err
	}
	url := "/backend-api/conversation/" + conversationID
	headers := map[string]string{
		"Authorization": "Bearer " + aiTokenModel.Token,
		"Content-Type":  "application/json; charset=utf-8",
		"User-Agent":    "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Referer":       "https://chat.openai.com",
		"Origin":        "https://chat.openai.com",
		"Cache-Control": "no-cache",
	}
	result, err := SendRequest("web", "PATCH", url, headers, body)
	if err != nil {
		return "", err
	}
	return result, nil
}
