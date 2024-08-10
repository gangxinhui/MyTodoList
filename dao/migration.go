package dao

import (
	"MyTodoList/model"
)

// 执行数据迁移
func migration() {
	// 自动迁移模式
	/*
			AutoMigrate函数旨在使数据库模式与代码中定义结构体保持同步，当你的User结构体（或其他结构体）的结构更新时，调用AutoMigrate将会相应地更新数据库表。
				创建表：如果数据库中不存在模型对应的表，AutoMigrate会创建它。
				添加新列：如果在结构体中添加了新字段，AutoMigrate会在表中添加相应的列。
				修改列类型：如果更改了字段的类型，AutoMigrate可以在表中更新该列的类型，但有一些限制。例如，它不能自动将一种类型的列更改为另一种类型，如果这需要数据转换或存在数据丢失的风险。
				创建索引：如果在结构体字段中添加了索引注释，AutoMigrate会在表中创建相应的索引。
		"gorm:table_options", "charset=utf8mb4"就是创建这个新表的时候使用 UTF-8 字符集
	*/
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{}, &model.Task{})
	if err != nil {
		return
	}
	// DB.Model(&Task{}).AddForeignKey("uid","User(id)","CASCADE","CASCADE")
}
