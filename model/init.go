package model

import (
	"fmt"

	"dhui.com/configs"
	"dhui.com/funcs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB = Init()

func Init() *gorm.DB {
	// dsn := "host=123.60.161.44 user=pg password=pg dbname=My_Project port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		configs.SQL_USER, configs.SQL_PASSWORD, configs.SQL_HOST, configs.SQL_PORT, configs.DATABASE)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		funcs.Danger("数据库连接失败! ", err)
		// panic("数据库连接失败")
	}
	return db
}
