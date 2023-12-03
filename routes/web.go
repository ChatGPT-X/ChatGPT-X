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
		// 用户管理
		userGroup := apiGroup.Group("user")
		{
			controller := controllers.UserController{}
			userGroup.POST("doRegister", controller.DoRegister) // 用户注册
			userGroup.POST("doLogin", controller.DoLogin)       // 用户登录
		}
		// AI 模型关系映射
		aiModelMapGroup := apiGroup.Group("aiModelMap").Use(middlewares.CheckAdmin())
		{
			controller := controllers.AiModelMapController{}
			aiModelMapGroup.POST("create", controller.Create) // 创建 AI 模型关系映射
			aiModelMapGroup.PUT("update", controller.Update)  // 修改 AI 模型关系映射
			aiModelMapGroup.GET("select", controller.Select)  // 查询 AI 模型关系映射列表
		}
	}
}
