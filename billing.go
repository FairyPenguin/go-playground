package main

import "fmt"

func bill() {
	var pl = fmt.Println

	pl(calculateDiscountRate(7421))
}

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	// ?
	totalCostBeforeDiscount := calculateBaseBill(costPerMessage, numMessages)

	discountValue := calculateDiscountRate(numMessages)

	totalCostAfterDiscount := totalCostBeforeDiscount - (totalCostBeforeDiscount * discountValue)

	return totalCostAfterDiscount
}

func calculateDiscountRate(messagesSent int) float64 {
	/*
	   10% for more than 1000 messages.
	   20% for more than 5000 messages.
	   0% for anything less.
	*/

	var discountRate float64

	if messagesSent > 1000 && messagesSent < 5000 {

		discountRate = 0.1

	} else if messagesSent > 5000 {

		discountRate = 0.2

	} else {
		discountRate = 0
	}

	return discountRate
}

// don't touch below this line

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}
