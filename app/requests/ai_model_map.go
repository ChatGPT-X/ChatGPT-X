package requests

// ValidateAiModelMapCreate 验证 AI 模型关系映射表创建验证。
type ValidateAiModelMapCreate struct {
	DisplayName string `form:"display_name" binding:"required,min=2,max=30"`
	AiName      string `form:"ai_name" binding:"required,min=2,max=30"`
	Status      uint   `form:"status" binding:"oneof=0 1"`
}

// ValidateAiModelMapUpdate 验证 AI 模型关系映射表更新验证。
type ValidateAiModelMapUpdate struct {
	ID          uint   `form:"id" binding:"required,numeric"`
	DisplayName string `form:"display_name" binding:"required,min=2,max=30"`
	AiName      string `form:"ai_name" binding:"required,min=2,max=30"`
	Status      uint   `form:"status" binding:"oneof=0 1"`
}

// ValidateAiModelMapSelect 验证 AI 模型关系映射表查询验证。
type ValidateAiModelMapSelect struct {
	Page     int64 `form:"page" binding:"numeric"`
	PageSize int64 `form:"page_size" binding:"numeric"`
}

// ValidateAiModelMapDelete 验证 AI 模型关系映射表删除验证。
type ValidateAiModelMapDelete struct {
	ID uint `form:"id" binding:"required,numeric"`
}
