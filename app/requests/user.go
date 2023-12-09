package requests

// ValidateDoRegister 用户注册验证。
type ValidateDoRegister struct {
	Username string `form:"username" binding:"required,alphanum,min=4,max=30"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=6,max=30"`
}

// ValidateDoLogin 用户登录验证。
type ValidateDoLogin struct {
	Username string `form:"username" binding:"required,alphanum,min=4,max=30"`
	Password string `form:"password" binding:"required,min=6,max=30"`
}

// ValidateUserList 用户列表验证。
type ValidateUserList struct {
	Page     int64 `form:"page" binding:"numeric"`
	PageSize int64 `form:"page_size" binding:"numeric"`
}

// ValidateUserUpdate 用户更新验证。
type ValidateUserUpdate struct {
	ID        uint   `form:"id" binding:"required,numeric"`
	AiTokenID *uint  `form:"ai_token_id" binding:"omitnil,numeric"`
	Password  string `form:"password" binding:"omitempty,min=6,max=30"`
	Status    string `form:"status" binding:"required,oneof=y n"`
}

// ValidateUserDelete 用户删除验证。
type ValidateUserDelete struct {
	ID uint `form:"id" binding:"required,numeric"`
}
