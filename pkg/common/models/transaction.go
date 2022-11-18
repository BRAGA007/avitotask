package models

import "time"

type Transaction struct {
	UserId      int `json:"user_id"`
	CreatedAt   time.Time
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}
