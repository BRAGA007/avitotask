package models

import "time"

type Revenue struct {
	UserId    int `json:"user_id"`
	ServiceId int `json:"service_id"`
	OrderId   int `gorm:"primary_key"`
	Amount    int `json:"amount"`
	CreatedAt time.Time
}
