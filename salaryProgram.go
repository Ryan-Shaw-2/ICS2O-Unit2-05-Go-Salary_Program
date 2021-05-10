// Created by: Ryan-Shaw-2
// Created on: May 2021
//
// This program calculates the users salary

package main

import (
	"fmt"

	"github.com/leekchan/accounting"
)

func main() {
	// This function displays currency
	accountingFormater := accounting.Accounting{Symbol: "$", Precision: 2}
	var hoursWorked int
	var hourlyRate int

	// input
	fmt.Println("This program calculates the users salary")
	fmt.Println()
	fmt.Print("Enter your number of hours worked: ")
	fmt.Scanln(&hoursWorked)
	fmt.Print("Enter your hourlyRate: ")
	fmt.Scanln(&hourlyRate)
	fmt.Println()

	// process
	var pay = float64(hoursWorked*hourlyRate) * (1.00 - 0.18)
	var tax = 0.18 * float64(hoursWorked*hourlyRate)

	// output
	fmt.Println("Your pay will be:", accountingFormater.FormatMoney(pay))
	fmt.Println("You will pay", accountingFormater.FormatMoney(tax), "in taxes")
}
