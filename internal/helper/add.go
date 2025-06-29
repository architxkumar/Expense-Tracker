package helper

import (
	"Expense-Tracker/internal/model"
	"encoding/json"
	"errors"
	"io"
	"os"
)

func AddExpense(description string, amount int, file *os.File) error {
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if description == "" {
		return errors.New("description must not be empty")
	}
	expenseEntry := model.Expense{
		Description: description,
		Amount:      amount,
	}
	fileContents, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	var expenseList []model.Expense
	if len(fileContents) != 0 {
		err = json.Unmarshal(fileContents, &expenseList)
		if err != nil {
			return err
		}

		// To overwrite file contents
		err = os.Truncate(file.Name(), 0)
		if err != nil {
			return err
		}
		// To place file pointer to start before writing to file
		_, err = file.Seek(0, 0)
		if err != nil {
			return err
		}
	}
	expenseList = append(expenseList, expenseEntry)
	output, err := json.Marshal(expenseList)
	if err != nil {
		return err
	}
	_, err = file.Write(output)
	if err != nil {
		return err
	}
	return nil
}
