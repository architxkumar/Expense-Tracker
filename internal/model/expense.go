package model

type Expense struct {
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}
