package middlewares

import (
	"chatgpt_x/app/models/user"
	"chatgpt_x/pkg/app"
	"chatgpt_x/pkg/e"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CheckAdmin 验证是否为管理员。
func CheckAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		isAdmin, ok := session.Get("user_is_admin").(uint)
		if !ok || isAdmin != user.IsAdmin {
			c.AbortWithStatusJSON(http.StatusOK, app.Response{
				Code: e.ErrorAuthFail,
				Msg:  e.GetMsg(e.ErrorAuthFail),
				Data: nil,
			})
		}
		c.Next()
	}
}

// CheckLogin 验证是否登录。
func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("user_id") == nil {
			c.AbortWithStatusJSON(http.StatusOK, app.Response{
				Code: e.ErrorAuthFail,
				Msg:  e.GetMsg(e.ErrorAuthFail),
				Data: nil,
			})
		}
		c.Next()
	}
}
