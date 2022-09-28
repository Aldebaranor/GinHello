package model

import "time"

type User_info struct {
	UId        string    `json:"u_id"`
	UserName   string    `json:"user_name"`
	CreateTime time.Time `json:"create_time"`
}
