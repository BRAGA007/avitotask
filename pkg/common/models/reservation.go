package models

type Reservation struct {
	User_Id    uint   `json:"user_id"`
	Service_Id uint   `json:"service_id"`
	Order_Id   uint   `gorm:"primary_key"`
	Cost       int    `json:"cost"`
	Status     string `json:"status"`
}
