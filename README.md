# Expense Tracker

Expense Tracker is a simple CLI based tool to store and manage your daily expenses.
<br/>
Built as a learning project for [roadmap.sh](https://roadmap.sh/projects/expense-tracker) challenge in Golang using 
[unfave/cli](https://github.com/urfave/cli) framework.
<br/>

## Features

- Add expense with amount and description
- View all expenses
- Delete specific expense record 
- Expense summarization for all records or by month
- Store data locally in JSON file
- Export contents to CSV file

## Usage
Clone the Repo and compile the program:
```bash
git clone https://github.com/architxkumar/Expense-Tracker.git
cd Expense-Tracker
go build -o expense-tracker
```
Command usage:
```bash
# Adding an Expense
./expense-tracker add --description "Lunch" --amount 20
./expense-tracker add --description "Dinner" --amount 10

# Viewing an Expense Records
./expense-tracker list

# Deleting an Expense Record
./expense-tracker delete --id 2

# Expense summary 
./expense-tracker summary

# Expense summary for a month
./expense-tracker summary --month 1

# Exporting contents to csv file
./expense-tracker export

# Command usage
./expense-tracker --help
```

## Licence

This project is licensed under the [MIT License](./LICENSE)
