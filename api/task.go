package api

import (
	"MyTodoList/pkg/util"
	"MyTodoList/service"
	"MyTodoList/types"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.CreateTaskReq
		if err := ctx.ShouldBind(&req); err == nil {
			//参数校验
			l := service.GetTaskSrv()
			resp, err := l.CreateTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

	}

}
func ListTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.TaskList
		//ShouldBind 要是传递的参数的key 不一样 会不检测直接err=nil
		if err := ctx.ShouldBind(&req); err == nil {
			//参数校验
			if req.Limit == 0 {
				req.Limit = 15
			}
			l := service.GetTaskSrv()
			resp, err := l.ListTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			fmt.Println("错误")
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}
	}
}
func ShowListTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.ShowTaskReq
		if err := ctx.ShouldBindQuery(&req); err == nil {
			//参数校验
			l := service.GetTaskSrv()
			resp, err := l.ShowTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))

		}

	}

}
func UpdateTaskHandler() gin.HandlerFunc {
	/*
			req := types.UpdateTaskReq跟req := new(types.UpdateTaskReq)的区别？
			   1·在这段代码中，req := types.UpdateTaskReq 仅仅是声明了一个类型为 types.UpdateTaskReq 的变量，
					但没有实际创建该类型的实例。ctx.ShouldBind(&req) 会失败，因为它期待的是一个实例而不是类型。
				2.在这段代码中，req := new(types.UpdateTaskReq) 创建了一个 types.UpdateTaskReq 类型的实例，
		       ctx.ShouldBind(req) 可以正确地将请求中的数据绑定到该实例上。之后，调用 l.UpdateTask(ctx.Request.Context(), req)
		也能够正确地传递该实例，并进行相应的数据库更新操作。

	*/
	return func(ctx *gin.Context) {
		req := new(types.UpdateTaskReq)
		if err := ctx.ShouldBind(&req); err == nil {
			//参数校验
			l := service.GetTaskSrv()
			resp, err := l.UpdateTask(ctx.Request.Context(), req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

	}

}
func SearchTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.SearchTaskReq
		if err := ctx.ShouldBind(&req); err == nil {
			l := service.GetTaskSrv()
			resp, err := l.SearchTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

	}

}
func DeleteTaskHander() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.DeleteTaskReq
		if err := ctx.ShouldBind(&req); err == nil {
			l := service.GetTaskSrv()
			resp, err := l.DeleteTask(ctx.Request.Context(), &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
				return
			}
			ctx.JSON(http.StatusOK, resp)
		} else {
			util.LogrusObj.Infoln(err)
			ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		}

	}

}
