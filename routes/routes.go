package routes

import (
	"MyTodoList/api"
	"MyTodoList/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors()) // 解决跨域问题
	v1 := r.Group("api/v1")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		v1.POST("user/register", api.UserRegister())  // 用户注册
		v1.POST("user/login", api.UserLoginHandler()) // 用户登陆
		//任务操作
		authed := v1.Group("/")
		authed.Use(middleware.JWT())
		{
			authed.POST("task_create", api.CreateTaskHandler()) //创建任务
			authed.GET("task_list", api.ListTaskHandler())      // 获取用户任务列表
			authed.GET("task_show", api.ShowListTaskHandler())  //获取某条备忘录详情
			authed.POST("task_update", api.UpdateTaskHandler()) //修改某条备忘录信息
			authed.POST("task_search", api.SearchTaskHandler()) // 搜索用户备忘录
			authed.DELETE("task_delete", api.DeleteTaskHander())
		}
	}
	return r
}
