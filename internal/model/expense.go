package model

import "time"

type Expense struct {
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
	Time        time.Time `json:"time"`
}
