package api

import (
	"MyTodoList/pkg/util"
	"MyTodoList/service"
	"MyTodoList/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserServiceReq
		if err := ctx.ShouldBind(&req); err == nil {
			// 参数校验 确保在整个应用程序中只创建并使用一个 UserSrv 实例
			l := service.GetUserSrv()
			resp, err := l.Register(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			//ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}

}
func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.UserServiceReq
		if err := ctx.ShouldBind(&req); err == nil {
			l := service.GetUserSrv()
			resp, err := l.Login(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)

		}

	}

}
