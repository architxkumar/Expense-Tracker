package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Expense-Tracker"
	app.Usage = "A simple expense tracker to manage your finances."
	app.Commands = []cli.Command{
		{Name: "add", Usage: "Add an expense with a description and amount", Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "description, d",
				Usage:    "Description of the expense",
				Required: true,
			},
		}, Action: func(c *cli.Context) error {
			description := c.String("description")
			fmt.Printf("Adding expense to %s\n", description)
			return nil
		},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
