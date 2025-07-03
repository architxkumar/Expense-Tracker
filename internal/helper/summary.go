package helper

import (
	"Expense-Tracker/internal/model"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const AllMonths = 0

// ExpenseSummary reads the contents from the file and
// prints the total expense for a given month.
// If supplied with month = 0, it will print the summary
// for all records. It will return an error in case of error reading
// file contents or parsing JSON
func ExpenseSummary(file *os.File, month int) error {
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
				// 0 represents non flag invocation command usage
				if month == AllMonths {
					sum += expense.Amount
				} else {
					if int(expense.Time.Month()) == month {
						sum += expense.Amount
					}
				}
			}
			if sum != 0 {
				fmt.Printf("Total Expenses: $%d\n", sum)
			} else {
				fmt.Printf("No records for the supplied month\n")
			}
		} else {
			fmt.Println("No expense found")
		}
	} else {
		fmt.Println("File Contents Empty")
	}
	return nil
}
