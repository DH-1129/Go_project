package model

import "time"

type Comment struct {
	Id          int       `gorm:"column:id;primary_key;AUTO_INCREMENT;omitempty"`
	Uid         string    `gorm:"column:uid;omitempty"`
	V_id        int       `gorm:"column:v_id;omitempty"`
	Reply_Id    int       `gorm:"column:reply_id;omitempty"`
	Create_Time time.Time `gorm:"column:create_time;omitempty"`
	Content     string    `gorm:"column:content;omitempty"`
}

func (c Comment) TableName() string {
	return "comments"
}
