package main

import (
	"fmt"
)

//Expenses will have a name and an amount
type Expenses struct {
	name   string
	amount float64
}

func main() {
	todayexpenses := Expenses{
		name:   "meal",
		amount: 12.5,
	}

	yesterdayexpenses := Expenses{
		name:   "present",
		amount: 20.0,
	}

	totalexpenses := todayexpenses.amount + yesterdayexpenses.amount

	fmt.Println("Today you have spent:", todayexpenses.amount, "in", todayexpenses.name)
	fmt.Println("Yesterday you spent:", yesterdayexpenses.amount, "in", yesterdayexpenses.name)
	fmt.Println("Total spent this month:", totalexpenses)
}
