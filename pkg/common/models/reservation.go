package models

type Reservation struct {
	User_Id    int    `json:"user_id"`
	Service_Id int    `json:"service_id"`
	Order_Id   int    `gorm:"primary_key"`
	Cost       int    `json:"cost"`
	Status     string `json:"status"`
}
