package databases

import (
	"fmt"

	"dhui.com/configs"
	"dhui.com/tools"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB = Init()

func Init() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		configs.SQL_USER, configs.SQL_PASSWORD, configs.SQL_HOST, configs.SQL_HOST, configs.DATABASE)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		tools.Danger(err, " 数据库连接失败!")
	}
	return db
}
