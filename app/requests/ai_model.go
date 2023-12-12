package requests

// ValidateAiModelCreate AI 模型创建验证。
type ValidateAiModelCreate struct {
	Type      uint   `form:"type" binding:"required,oneof=1 2"`
	AliasName string `form:"alias_name" binding:"required,min=1,max=30"`
	Name      string `form:"name" binding:"required,min=1,max=30"`
	Status    string `form:"status" binding:"required,oneof=y n"`
}

// ValidateAiModelUpdate AI 模型更新验证。
type ValidateAiModelUpdate struct {
	ID        uint   `form:"id" binding:"required,numeric"`
	Type      uint   `form:"type" binding:"required,oneof=1 2"`
	AliasName string `form:"alias_name" binding:"required,min=1,max=30"`
	Name      string `form:"name" binding:"required,min=1,max=30"`
	Status    string `form:"status" binding:"required,oneof=y n"`
}

// ValidateAiModelList AI 模型查询列表验证。
type ValidateAiModelList struct {
	Page     int64 `form:"page" binding:"numeric"`
	PageSize int64 `form:"page_size" binding:"numeric"`
}

// ValidateAiModelDelete AI 模型删除验证。
type ValidateAiModelDelete struct {
	ID uint `form:"id" binding:"required,numeric"`
}
