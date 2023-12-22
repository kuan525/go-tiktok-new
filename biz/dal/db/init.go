package db

import (
	"go-tiktok-new/pkg/constants"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var DB *gorm.DB

func Init() {
	var err error

	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN),
		&gorm.Config{
			// 预准备语句，先把类似于带？的语句编译，然后直接执行，可复用执行，同时可以防止sql注入
			PrepareStmt: true,
			// 忽略默认开启事务行为
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	// 分布式追踪开放标准，统一服务以服务之间访问的规范，想要实际查看trace，需要接入一个追踪系统，例如jaeger等
	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}
}
