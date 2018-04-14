package main

import (
	"fmt"

	"github.com/nfuste/learning/expenses"
)

func main() {

	todayexpenses := expenses.Expenses{
		Name:   "meal",
		Amount: 12.5,
	}

	yesterdayexpenses := expenses.Expenses{
		Name:   "present",
		Amount: 20.0,
	}

	totalexpenses := todayexpenses.Amount + yesterdayexpenses.Amount

	fmt.Println("Today you have spent:", todayexpenses.Amount, "in", todayexpenses.Name)
	fmt.Println("Yesterday you spent:", yesterdayexpenses.Amount, "in", yesterdayexpenses.Name)
	fmt.Println("Total spent this month:", totalexpenses)
}
