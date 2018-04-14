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

	fmt.Println("Today you have spent:", todayexpenses.amount, "in", todayexpenses.name)
}
