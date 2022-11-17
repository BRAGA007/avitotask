package models

import "time"

type Revenue struct {
	User_Id    int `json:"user_id"`
	Service_Id int `json:"service_id"`
	Order_Id   int `gorm:"primary_key"`
	Amount     int `json:"amount"`
	CreatedAt  time.Time
}
