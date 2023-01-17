package model

import (
	"time"
)

type User struct {
	// gorm.Model
	Id          int       `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Uid         string    `gorm:"column:uid"`
	Username    string    `gorm:"column:username"`
	Email       string    `gorm:"column:email"`
	Phone       string    `gorm:"column:phone"`
	Img_url     string    `gorm:"column:img_url"`
	Sex         int       `gorm:"column:sex"`
	Signature   string    `gorm:"column:signature"`
	Create_time time.Time `gorm:"column:create_time"`
	Last_time   time.Time `gorm:"column:last_time"`
	Password    string    `gorm:"column:password"`
}

func (u User) TableName() string {
	return "user_info"
}

func GetEmail(email string) (u User, err error) {
	var user User
	err = DB.Where("email=?", email).First(&user).Error
	return user, err
}
