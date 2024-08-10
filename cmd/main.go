package main

import (
	conf "MyTodoList/config"
	"MyTodoList/dao"
	"MyTodoList/repository/cache"
	"MyTodoList/routes"
)

func main() {
	loading()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
func loading() {
	conf.Init()
	dao.MySQLInit()
	cache.Redis()

}
