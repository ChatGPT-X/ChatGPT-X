package requests

// ValidateTokenCreate 密钥创建验证。
type ValidateTokenCreate struct {
	Type   uint   `form:"type" binding:"required,oneof=0 1"`
	Token  string `form:"token" binding:"required,min=1,max=2000"`
	Remark string `form:"remark" binding:"min=1,max=255"`
	Status uint   `form:"status" binding:"required,oneof=0 1"`
}

// ValidateTokenUpdate 密钥更新验证。
type ValidateTokenUpdate struct {
	ID     uint   `form:"id" binding:"required,numeric"`
	Type   uint   `form:"type" binding:"required,oneof=0 1"`
	Token  string `form:"token" binding:"required,min=1,max=2000"`
	Remark string `form:"remark" binding:"min=1,max=255"`
	Status uint   `form:"status" binding:"required,oneof=0 1"`
}

// ValidateTokenList 密钥查询列表验证。
type ValidateTokenList struct {
	Page     int64 `form:"page" binding:"numeric"`
	PageSize int64 `form:"page_size" binding:"numeric"`
}

// ValidateTokenDelete 密钥删除验证。
type ValidateTokenDelete struct {
	ID uint `form:"id" binding:"required,numeric"`
}
