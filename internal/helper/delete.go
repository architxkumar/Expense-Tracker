package helper

import (
	"Expense-Tracker/internal/model"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

// DeleteTask reads contents from the file and deletes
// an expense entry with the associated id. It returns an error
// in case of reading file contents
func DeleteTask(file *os.File, id int) error {
	// The task id supplied by the user has been passed
	// from the perspective that the index starts from 1
	if id < 1 {
		return errors.New("id must be greater than zero")
	}
	fileContents, err := io.ReadAll(file)
	if err == nil {
		if len(fileContents) > 0 {
			var expenseList []model.Expense
			err = json.Unmarshal(fileContents, &expenseList)
			if err != nil {
				return err
			}
			length := len(expenseList)
			expenseIndex := id - 1
			if expenseIndex < length {
				expenseList = append(expenseList[:expenseIndex], expenseList[expenseIndex+1:]...)
				output, err := json.Marshal(expenseList)
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
				_, err = file.Write(output)
				if err != nil {
					return err
				}

			} else {
				fmt.Println("No tasks with the specified id exists")
				return nil
			}

		} else {
			fmt.Println("No records exist")
			return nil
		}
	} else {
		return errors.New("error reading file content")
	}
	return nil
}
