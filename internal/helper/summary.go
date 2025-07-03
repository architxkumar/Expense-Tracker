package helper

import (
	"Expense-Tracker/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ExpenseSummary(file *os.File) error {
	fileContents, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	if len(fileContents) > 0 {
		var expenseList []model.Expense
		err = json.Unmarshal(fileContents, &expenseList)
		if err != nil {
			return err
		}
		if len(expenseList) > 0 {
			var sum int
			for _, expense := range expenseList {
				sum += expense.Amount
			}
			fmt.Printf("Total Expenses: $%d\n", sum)
		} else {
			fmt.Println("No expense found")
		}
	} else {
		fmt.Println("File Contents Empty")
	}
	return nil
}
