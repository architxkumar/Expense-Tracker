package main

import (
	"Expense-Tracker/internal/helper"
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
			fmt.Println("Adding expense ", amount, " with description:", description)
			err := helper.AddExpense(description, amount, file)
			return err
		},
			UsageText: "add [Flag] <Value>",
			CustomHelpTemplate: `NAME:
        {{.HelpName}} - {{.Usage}}

USAGE: 
        {{.HelpName}} -d <description> -a <amount>

MANDATORY FLAGS:
        -d, --description   Description of the expense
        -a, --amount        Amount of the expense
`,
		},
	}
	err = app.Run(os.Args)
	if err != nil {
		log.Fatal("Error initializing application: ", err.Error())
	}
}
