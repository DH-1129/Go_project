package model

import (
	"time"
)

type User struct {
	// gorm.Model
	Id          int    `gorm:"column:id;primary_key;AUTO_INCREMENT;omitempty"`
	Uid         string `gorm:"column:uid;omitempty"`
	Username    string `gorm:"column:username;omitempty" `
	Email       string `gorm:"column:email;omitempty" `
	Phone       string `gorm:"column:phone;omitempty" `
	Img_Name    string `gorm:"column:img_name;omitempty"`
	Sex         string `gorm:"column:sex;omitempty"`
	Signature   string `gorm:"column:signature;omitempty"`
	Create_Time time.Time
	Update_Time time.Time `gorm:"autoUpdateTime"`
	Password    string    `gorm:"column:password;omitempty"`
	Likes       string    `gorm:"column:likes"`
	Collections string    `gorm:"column:collections"`
	Prex        string    `gorm:"column:prex"`
}

func (u User) TableName() string {
	return "user_info"
}

// func (u User) Value() (driver.Value, error) {
// 	user, err := json.Marshal(u)
// 	return string(user), err
// }
// func (u *User) Scan(input interface{}) error {
// 	return json.Unmarshal(input.([]byte), u)
// }

func GetEmail(email string) (u User, err error) {
	var user User
	err = DB.Where("email=?", email).First(&user).Error
	return user, err
}
