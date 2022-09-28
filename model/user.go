package model

import "time"

func (User_info) TableName() string {
	return "user_info"
}

type User_info struct {
	UId        string    `json:"u_id"`
	UserName   string    `json:"user_name"`
	CreateTime time.Time `json:"create_time"`
}
