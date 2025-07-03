package helper

import (
	"Expense-Tracker/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// ListTask reads expense records from file and prints them
// in a table. It prints "No records found" in case of empty file.
// It returns an error in case of file opening or invalid JSON.
func ListTask(file *os.File) error {
	fileContents, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	if len(fileContents) == 0 {
		fmt.Println("No records exist")
		return nil
	} else {
		var expenseList []model.Expense
		// When the json array is present inside "[]"
		// but is empty
		if len(expenseList) == 0 {
			fmt.Println("No records exist")
			return nil
		}
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
