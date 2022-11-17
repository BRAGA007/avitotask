package models

type User struct {
	ID      int `gorm:"primary_key"`
	Balance int `json:"balance"`
}
