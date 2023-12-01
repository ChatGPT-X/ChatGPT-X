package routes

import (
	"chatgpt_x/app/http/controllers"
	"github.com/gin-gonic/gin"
)

// Register used for register router.
func Register(engine *gin.Engine) {
	apiGroup := engine.Group("api")
	{
		userGroup := apiGroup.Group("user")
		{
			controller := controllers.UserController{}
			userGroup.POST("doRegister", controller.DoRegister) // 用户注册
			userGroup.POST("doLogin", controller.DoLogin)       // 用户登录
		}
	}
}
