package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {

	bill()

	fmt.Println("GOGO")
	// _ = monthlyBillIncrease
	// _ = getBillForMonth
	printReports(
		"Welcome to the Hotel California",
		"Such a lovely place",
		"Plenty of room at the Hotel California",
	)
}

func printCostReport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}

func printReports(intro, body, outro string) {

	printCostReport(func(message string) int {
		messageLen := len(message)
		return messageLen * 2
	}, intro)

	printCostReport(func(message string) int {
		messageLen := len(message)
		return messageLen * 3
	}, body)

	printCostReport(func(message string) int {
		messageLen := len(message)
		return messageLen * 4
	}, outro)

}

// func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
// 	var lastMonthBill int = getBillForMonth(costPerSend, numLastMonth)
// 	var thisMonthBill int = getBillForMonth(costPerSend, numThisMonth)
// 	pl(thisMonthBill - lastMonthBill)
// 	return thisMonthBill - lastMonthBill
// }

// func getBillForMonth(costPerSend, messagesSent int) int {
// 	var bill int = costPerSend * messagesSent

// 	return bill
// }

// func getMonthlyPrice(tier string) int {
// 	var monthlyPrice int

// 	if tier == "basic" {
// 		monthlyPrice = 100.00 * 100
// 		pl(monthlyPrice)

// 		return monthlyPrice
// 	} else if tier == "premium" {
// 		monthlyPrice = 150.00 * 100
// 		return monthlyPrice

// 	} else if tier == "enterprise" {
// 		monthlyPrice = 500.00 * 100
// 		return monthlyPrice

// 	} else {
// 		monthlyPrice = 0
// 		return monthlyPrice
// 	}

// }

/*

Variables in Go are passed by value (except for a few data types we haven't covered yet).
"Pass by value" means that when a variable is passed into a function,
that function receives a copy of the variable.
The function is unable to mutate the caller's original data.



*/
