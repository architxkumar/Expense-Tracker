package model

import "time"

// Expense represents an expenditure entry in the file
type Expense struct {
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	// Time is automatically generated at expense insertion
	Time time.Time `json:"time"`
}
