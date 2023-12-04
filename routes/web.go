package routes

import (
	"chatgpt_x/app/http/controllers"
	"chatgpt_x/app/http/middlewares"
	"github.com/gin-gonic/gin"
)

// Register used for register router.
func Register(engine *gin.Engine) {
	apiGroup := engine.Group("api")
	{
		// 免登录接口
		controller := controllers.UserController{}
		apiGroup.POST("doRegister", controller.DoRegister) // 用户注册
		apiGroup.POST("doLogin", controller.DoLogin)       // 用户登录
		apiGroup.GET("logout", controller.Logout)          // 用户登出
		// 前台接口
		//userGroup := apiGroup.Group("user", middlewares.CheckLogin())
		//{
		//
		//}

		// 后台接口
		adminGroup := apiGroup.Group("admin", middlewares.CheckAdmin())
		{
			// 用户管理
			userGroup := adminGroup.Group("user")
			{
				controller := controllers.UserController{}
				userGroup.DELETE("delete", controller.Delete) // 删除用户
			}
			// AI 模型管理
			aiModelGroup := adminGroup.Group("aiModel")
			{
				controller := controllers.AiModelController{}
				aiModelGroup.POST("create", controller.Create)   // 创建 AI 模型
				aiModelGroup.PUT("update", controller.Update)    // 修改 AI 模型
				aiModelGroup.GET("list", controller.List)        // 查询 AI 模型列表
				aiModelGroup.DELETE("delete", controller.Delete) // 删除 AI 模型
			}
		}
	}
}
