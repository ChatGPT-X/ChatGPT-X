package requests

// ValidateUUIDv4 AI 密钥创建验证。
type ValidateUUIDv4 struct {
	ConversationID string `uri:"conversation_id" binding:"required,uuid4"`
}
