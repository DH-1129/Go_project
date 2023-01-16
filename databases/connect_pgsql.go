package databases

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Session_pg() (*gorm.DB, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("error:%v", err)
		}
	}()
	dsn := "host=123.60.161.44 user=pg password=pg dbname=My_Project port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	//dsn := "pg:pg@tcp(123.60.161.44:5432)/My_Project?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("pgsql连接失败: %v", err)
		return nil, err
	}
	return db, nil
}
