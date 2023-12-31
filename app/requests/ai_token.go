package requests

// ValidateAiTokenCreate AI 密钥创建验证。
type ValidateAiTokenCreate struct {
	Type   string `form:"type" binding:"required,oneof=access_token api_key"`
	Token  string `form:"token" binding:"required,min=1,max=2000"`
	Remark string `form:"remark" binding:"min=1,max=255"`
	Status string `form:"status" binding:"required,oneof=y n"`
}

// ValidateAiTokenUpdate AI 密钥更新验证。
type ValidateAiTokenUpdate struct {
	ID     uint   `form:"id" binding:"required,numeric"`
	Type   string `form:"type" binding:"required,oneof=access_token api_key"`
	Token  string `form:"token" binding:"required,min=1,max=2000"`
	Remark string `form:"remark" binding:"min=1,max=255"`
	Status string `form:"status" binding:"required,oneof=y n"`
}

// ValidateAiTokenList AI 密钥获取列表验证。
type ValidateAiTokenList struct {
	Page     int64 `form:"page" binding:"numeric"`
	PageSize int64 `form:"page_size" binding:"numeric"`
}

// ValidateAiTokenDelete AI 密钥删除验证。
type ValidateAiTokenDelete struct {
	ID uint `form:"id" binding:"required,numeric"`
}
