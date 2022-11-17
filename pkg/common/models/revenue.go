package models

import "time"

type Revenue struct {
	User_Id    uint `json:"user_id"`
	Service_Id uint `json:"service_id"`
	Order_Id   uint `gorm:"primary_key"`
	Sum        int  `json:"sum"`
	CreatedAt  time.Time
}
