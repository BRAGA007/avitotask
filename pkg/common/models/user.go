package models

type User struct {
	ID      uint `gorm:"primary_key"`
	Balance int  `json:"balance"`
}
