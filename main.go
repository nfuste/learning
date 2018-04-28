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

	if totalexpenses > 10 {
		fmt.Println("Mmm... looks like you have spent a lot")
	} else {
		fmt.Println("Great job! You have spent little money")
	}
}
