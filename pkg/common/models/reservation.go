package models

type Reservation struct {
	UserId    int    `json:"user_id"`
	ServiceId int    `json:"service_id"`
	OrderId   int    `gorm:"primary_key"`
	Cost      int    `json:"cost"`
	Status    string `json:"status"`
}
