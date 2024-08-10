package dao

import (
	conf "MyTodoList/config"
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

var _db *gorm.DB

func MySQLInit() {
	conn := strings.Join([]string{conf.DbUser, ":", conf.DbPassWord, "@tcp(", conf.DbHost, ":", conf.DbPort, ")/", conf.DbName, "?charset=utf8&parseTime=true"}, "")
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       conn,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: ormLogger, // 打印日志
		//NamingStrategy作用：设置表名、列名等的命名策略
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表明不加s
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	/*
		连接池的好处
		使用连接池可以显著提高应用程序的性能和稳定性，主要表现在以下几个方面：
		减少连接建立的开销：连接池复用已经建立的连接，避免每次数据库操作都需要重新建立连接。
		控制并发：通过设置最大连接数，可以防止过多的并发连接导致数据库过载。
		优化资源使用：通过设置空闲连接数和连接的最大生命周期，可以更高效地使用数据库资源，防止资源泄漏和连接过期问题。
	*/
	sqlDB.SetMaxIdleConns(20)  // 设置连接池中空闲连接的最大数量
	sqlDB.SetMaxOpenConns(100) // 设置最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	//这行代码将新创建的数据库连接 db 赋值给全局变量 _db。这样，其他地方的代码可以通过 _db 访问和使用这个数据库连接
	_db = db
	//migration()
}
func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
