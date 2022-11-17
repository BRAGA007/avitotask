package models

import "time"

type Revenue struct {
	User_Id    int `json:"user_id"`
	Service_Id int `json:"service_id"`
	Order_Id   int `gorm:"primary_key"`
	Sum        int `json:"sum"`
	CreatedAt  time.Time
}
