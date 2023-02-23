package model

import "time"

type Video_Info struct {
	Id               int `gorm:"column:id;primary_key;AUTO_INCREMENT;omitempty"`
	Uid              string
	Title            string
	Content          string
	Create_Time      time.Time
	File_Name        string
	Tag              string
	Prefix           string
	Comment_Count    int
	Like_Count       int
	Collection_Count int
}

func (v Video_Info) TableName() string {
	return "video_info"
}
