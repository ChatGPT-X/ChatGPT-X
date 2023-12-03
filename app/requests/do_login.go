package requests

// ValidateDoLogin 验证用户登录。
type ValidateDoLogin struct {
	Username string `form:"username" binding:"required,alphanum,min=4,max=30"`
	Password string `form:"password" binding:"required,min=6,max=30"`
}
