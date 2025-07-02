package helper

import (
	"Expense-Tracker/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func PrintTask(file *os.File) error {
	fileContents, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	if len(fileContents) == 0 {
		fmt.Println("No records exist")
		return nil
	} else {
		var expenseList []model.Expense
		if err := json.Unmarshal(fileContents, &expenseList); err != nil {
			return err
		}
		fmt.Printf("# %-4s  %-9s  %-5s  %s\n", "Id", "Date", "Amount", "Description")
		for index, expense := range expenseList {
			// Task Index is incremented by 1 for ease of access
			fmt.Printf("# %-4d  %v $%-5d  %s\n", index+1, expense.Time.Format("02-01-2006"), expense.Amount, expense.Description)
		}
	}
	return nil
}
