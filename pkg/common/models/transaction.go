package models

import "time"

type Transaction struct {
	User_id     int `json:"user_id"`
	CreatedAt   time.Time
	Description string `json:"description"`
}
