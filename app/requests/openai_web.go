package requests

// ValidateUUIDv4 AI 密钥创建验证。
type ValidateUUIDv4 struct {
	ConversationID string `uri:"conversation_id" binding:"required,uuid4"`
}

// ValidateGetConversationHistory 获取对话历史验证。
type ValidateGetConversationHistory struct {
	Offset int `form:"offset" binding:"numeric"`
	Limit  int `form:"limit" binding:"numeric"`
}
