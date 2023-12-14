package requests

// ValidateSettingUpdate 更新系统设置表单验证。
type ValidateSettingUpdate struct {
	ID         uint   `form:"id" binding:"required,numeric"`
	ApiBaseUrl string `form:"api_base_url" binding:"required,url,min=1,max=255"`
	ApiProxy   string `form:"api_proxy" binding:"min=1,max=255"`
	ApiTimeout uint   `form:"api_timeout" binding:"required,numeric"`
	WebBaseUrl string `form:"web_base_url" binding:"required,url,min=1,max=255"`
	WebProxy   string `form:"web_proxy" binding:"min=1,max=255"`
	WebTimeout uint   `form:"web_timeout" binding:"required,numeric"`
}
