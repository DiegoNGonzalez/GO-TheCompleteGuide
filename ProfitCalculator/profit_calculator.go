package main

import "fmt"

func main() {

	var revenue, expenses, taxRate float64

	fmt.Print("Enter the revenue amount: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter the expenses amount: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax:= revenue-expenses
	earningsAfterTax:=earningsBeforeTax*(1-taxRate/100)
	ratio:=earningsBeforeTax/earningsAfterTax
	fmt.Println(earningsBeforeTax)
	fmt.Println(earningsAfterTax)
	fmt.Println(ratio)
}