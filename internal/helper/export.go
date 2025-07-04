package helper

import (
	"Expense-Tracker/internal/model"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

// ExportExpenses reads contents from file and exports contents
// into a new csv file. It returns an error in case of
// issue in reading/writing contents to file, unmarshalling JSON
// or creating new file.
func ExportExpenses(file *os.File) error {
	fileContents, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	if len(fileContents) == 0 {
		fmt.Println("No expenses to export")
		return nil
	}
	var expenseList []model.Expense
	err = json.Unmarshal(fileContents, &expenseList)
	if err != nil {
		return err
	}
	if len(expenseList) == 0 {
		fmt.Println("No expenses found")
		return nil
	}
	var csvContent [][]string
	csvContent = append(csvContent, []string{
		"id", "amount", "description",
	})
	for index, expense := range expenseList {
		csvContent = append(csvContent, []string{
			strconv.Itoa(index + 1),
			fmt.Sprintf("$%d", expense.Amount),
			expense.Description,
		})

	}
	csvFile, err := os.Create("expenses.csv")
	if err != nil {
		return err
	}
	writer := csv.NewWriter(csvFile)
	err = writer.WriteAll(csvContent)
	if err != nil {
		return err
	}
	fmt.Println("Expenses Exported Successfully")
	return nil
}
