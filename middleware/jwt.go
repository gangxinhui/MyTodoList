package middleware

import (
	"MyTodoList/pkg/ctl"
	"MyTodoList/pkg/e"
	"MyTodoList/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		code = e.SUCCESS
		token := c.GetHeader("Authorization")
		if token == "" {
			code = http.StatusBadRequest
			c.JSON(e.InvalidParams, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "缺少Token",
			})
			c.Abort()
			return
		}
		//验证用户token
		claims, err := util.ParseToken(token)
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		}
		if code != e.SUCCESS {
			c.JSON(e.InvalidParams, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "可能是身份过期了，请重新登录",
			})
			c.Abort()
			return
		}
		c.Request = c.Request.WithContext(ctl.NewContext(c.Request.Context(), &ctl.UserInfo{Id: claims.Id}))
		c.Next()

	}
}
