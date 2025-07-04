package main

import (
	"Expense-Tracker/internal/helper"
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

const filePath string = "data.json"

func main() {
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal("Error opening file")
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Panic("Error closing file: ", err.Error())
		}
	}()
	app := cli.NewApp()
	app.Name = "Expense-Tracker"
	app.Usage = "A simple expense tracker to manage your finances."
	app.Commands = []cli.Command{
		{Name: "add", Usage: "Add an expense with a description and amount", HelpName: "add", Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "d, description",
				Usage:    "Description of the expense",
				Required: true,
			},
			cli.IntFlag{
				Name:     "a, amount",
				Usage:    "Amount of the expense",
				Required: true,
			},
		}, Action: func(c *cli.Context) error {
			description := c.String("description")
			amount := c.Int("amount")
			fmt.Printf("Adding an expense of $%d with description: %s\n", amount, description)
			err := helper.AddExpense(description, amount, file)
			return err
		},
			CustomHelpTemplate: `NAME:
        {{.HelpName}} - {{.Usage}}

USAGE: 
        {{.HelpName}} -d <description> -a <amount>

MANDATORY FLAGS:
        -d, --description   Description of the expense
        -a, --amount        Amount of the expense
`,
		},
		{
			Name:  "list",
			Usage: "List all expenses",
			Action: func(c *cli.Context) error {
				err := helper.ListTask(file)
				return err
			},
		},
		{
			Name:  "delete",
			Usage: "Delete an expense",
			Action: func(c *cli.Context) error {
				id := c.Int("id")
				err := helper.DeleteTask(file, id)
				if err == nil {
					fmt.Println("Task Deleted Successfully")
				}
				return err
			},
			HelpName: "delete",
			CustomHelpTemplate: `NAME:
		{{.HelpName}} - {{.Usage}}
USAGE:
		{{.HelpName}} --id <id>

MANDATORY FLAG:
		--id			Id of the expense

`,
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:     " id",
					Usage:    "Expense ID",
					Required: true,
				},
			},
		},
		{
			Name:     "summary",
			Usage:    "Show summary of expenses",
			HelpName: "summary",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:     "month",
					Required: false,
					Usage:    "Month of expenses",
					Value:    0,
				},
			},
			Action: func(c *cli.Context) error {
				month := c.Int("month")
				// 0 represents non flag invocation use case
				if month < 0 || month > 12 {
					return errors.New("invalid month")
				}
				err := helper.ExpenseSummary(file, month)
				return err
			},
			CustomHelpTemplate: `NAME:
		{{.HelpName}} - {{.Usage}}
USAGE:
		{{.HelpName}} [FLAG] [VALUE]

OPTIONAL FLAG:
		--month <month>		month of expenses

EXAMPLE:
{{.HelpName}} --month 		To find summary of all expenses
{{.HelpName}} --month 1		To find summary of expense for the month of January
`,
		},
		{
			Name:  "export",
			Usage: "Export expenses to a csv file",
			Action: func(c *cli.Context) error {
				err := helper.ExportExpenses(file)
				return err
			},
		},
	}
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal("Error initializing application: ", err.Error())
	}
}
