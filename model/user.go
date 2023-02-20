package model

import (
	"time"
)

type User struct {
	// gorm.Model
	Id          int    `gorm:"column:id;primary_key;AUTO_INCREMENT;omitempty"`
	Uid         string `gorm:"column:uid;omitempty" json:"uid"`
	Username    string `gorm:"column:username;omitempty" json:"username"`
	Email       string `gorm:"column:email;omitempty" json:"email"`
	Phone       string `gorm:"column:phone;omitempty" json:"phone"`
	Img_url     string `gorm:"column:img_url;omitempty" json:"img_url"`
	Sex         string `gorm:"column:sex;omitempty" json:"sex"`
	Signature   string `gorm:"column:signature;omitempty" json:"signature"`
	Create_Time time.Time
	Update_Time time.Time `gorm:"autoUpdateTime"`
	Password    string    `gorm:"column:password;omitempty" json:"password,omitempty"`
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
