package openai_service

type WebService struct{}

func (s *WebService) Conversation(userID uint, body any) (<-chan []byte, error) {
	// 获取当前用户的 token
	aiTokenModel, err := GetAiTokenFromUser(userID)
	if err != nil {
		return nil, err
	}
	url := "/backend-api/conversation"
	headers := map[string]string{
		"User-Agent":    "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Referer":       "https://chat.openai.com",
		"Origin":        "https://chat.openai.com",
		"Content-Type":  "application/json",
		"Accept":        "text/event-stream",
		"Authorization": "Bearer " + aiTokenModel.Token,
	}
	ch, err := SendStreamRequest("web", "POST", url, headers, body)
	if err != nil {
		return nil, err
	}
	return ch, nil
}
