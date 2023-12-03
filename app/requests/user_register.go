package requests

// ValidateDoRegister 验证用户注册。
type ValidateDoRegister struct {
	Username string `form:"username" binding:"required,alphanum,min=4,max=30"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=6,max=30"`
}
