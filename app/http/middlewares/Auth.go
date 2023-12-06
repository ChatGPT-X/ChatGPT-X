package middlewares

import (
	"chatgpt_x/app/models/user"
	"chatgpt_x/pkg/app"
	"chatgpt_x/pkg/auth"
	"chatgpt_x/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CheckAdmin 验证是否为管理员。
func CheckAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 header 中获取 jwt token
		jwt, err := auth.GetTokenFromHeader(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, app.Response{
				Code: e.ErrorAuthFail,
				Msg:  e.GetMsg(e.ErrorAuthFail),
				Data: nil,
			})
			return
		}
		// 从 jwt token 中获取用户权限
		var claims *auth.Claims
		claims, err = auth.ParseToken(jwt)
		if err != nil || claims.IsAdmin != user.IsAdmin {
			c.AbortWithStatusJSON(http.StatusOK, app.Response{
				Code: e.ErrorAuthFail,
				Msg:  e.GetMsg(e.ErrorAuthFail),
				Data: nil,
			})
			return
		}
		c.Next()
	}
}

// CheckLogin 验证是否登录。
func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 header 中获取 jwt token
		jwt, err := auth.GetTokenFromHeader(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusOK, app.Response{
				Code: e.ErrorAuthFail,
				Msg:  e.GetMsg(e.ErrorAuthFail),
				Data: nil,
			})
			return
		}
		var claims *auth.Claims
		claims, err = auth.ParseToken(jwt)
		if err != nil || claims.UserID == 0 {
			c.AbortWithStatusJSON(http.StatusOK, app.Response{
				Code: e.ErrorAuthFail,
				Msg:  e.GetMsg(e.ErrorAuthFail),
				Data: nil,
			})
			return
		}
		c.Next()
	}
}
