package requests

// ValidateAiModelCreate 验证 AI 模型表创建验证。
type ValidateAiModelCreate struct {
	DisplayName string `form:"display_name" binding:"required,min=2,max=30"`
	AiName      string `form:"ai_name" binding:"required,min=2,max=30"`
	Status      uint   `form:"status" binding:"oneof=0 1"`
}

// ValidateAiModelUpdate 验证 AI 模型表更新验证。
type ValidateAiModelUpdate struct {
	ID          uint   `form:"id" binding:"required,numeric"`
	DisplayName string `form:"display_name" binding:"required,min=2,max=30"`
	AiName      string `form:"ai_name" binding:"required,min=2,max=30"`
	Status      uint   `form:"status" binding:"oneof=0 1"`
}

// ValidateAiModelSelect 验证 AI 模型表查询验证。
type ValidateAiModelSelect struct {
	Page     int64 `form:"page" binding:"numeric"`
	PageSize int64 `form:"page_size" binding:"numeric"`
}

// ValidateAiModelDelete 验证 AI 模型表删除验证。
type ValidateAiModelDelete struct {
	ID uint `form:"id" binding:"required,numeric"`
}
