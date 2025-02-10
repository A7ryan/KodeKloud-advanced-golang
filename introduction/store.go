package main

import "fmt"

// Declare the Expense struct here
type Expense struct {
	Name   string
	Amount float64
	Date   string
}

// Implement the Total method to calculate the total amount spent
func Total(expenses []Expense) float64 {
	var total float64
	for _, expense := range expenses {
		total += expense.Amount
	}
	return total
}

// Implement the getName method on the Expense struct here
func (e Expense) getName() string {
    return e.Name
}

func main() {
	expenses := []Expense{
		Expense{"Grocery", 50.0, "2022-01-01"},
		Expense{"Gas", 30.0, "2022-01-02"},
		Expense{"Restaurant", 40.0, "2022-01-03"},
	}

	fmt.Println(Total(expenses))
	fmt.Println(expenses[0].getName())
}
