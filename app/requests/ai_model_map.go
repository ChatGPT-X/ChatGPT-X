package requests

// ValidateAiModelMapCreate 验证 AI 模型关系映射表创建表单。
type ValidateAiModelMapCreate struct {
	DisplayName string `form:"display_name" binding:"required,min=2,max=30"`
	AiName      string `form:"ai_name" binding:"required,min=2,max=30"`
	IsDisabled  uint   `form:"is_disabled" binding:"oneof=0 1"`
}

// ValidateAiModelMapUpdate 验证 AI 模型关系映射表更新表单。
type ValidateAiModelMapUpdate struct {
	ID          uint   `form:"id" binding:"required,numeric"`
	DisplayName string `form:"display_name" binding:"required,min=2,max=30"`
	AiName      string `form:"ai_name" binding:"required,min=2,max=30"`
	IsDisabled  uint   `form:"is_disabled" binding:"oneof=0 1"`
}

// ValidateAiModelMapSelect 验证 AI 模型关系映射表查询表单。
type ValidateAiModelMapSelect struct {
	Page     int64 `form:"page" binding:"numeric"`
	PageSize int64 `form:"page_size" binding:"numeric"`
}

// ValidateAiModelMapDelete 验证 AI 模型关系映射表删除表单。
type ValidateAiModelMapDelete struct {
	ID uint `form:"id" binding:"required,numeric"`
}
