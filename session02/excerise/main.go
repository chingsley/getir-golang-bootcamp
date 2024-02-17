package main

import (
	"fmt"
	"strconv"
)

const (
	silverAcount = iota + 1
	goldAccount
	platinumAccount
	silverDiscount   = 0.9
	goldDiscount     = 0.8
	platinumDiscount = 0.7
)

func main() {
	var inputStr string
	fmt.Print("Enter 1, 2 or 3 for membershipt type: ")
	fmt.Scanln(&inputStr)
	membershipLevel, err := strconv.Atoi(inputStr)
	if err != nil {
		fmt.Printf("membershipLevel Atoi error: %s", err)
		return
	}

	if membershipLevel < 1 || membershipLevel > 3 {
		fmt.Println("Invalid input. Please enter 1, 2 or 3.")
		return
	}

	totalPrice := 0

	for i := 0; i < 3; i++ {
		fmt.Print("Enter a non-negative integer for the book price: ")
		fmt.Scanln(&inputStr)
		bookPrice, err := strconv.Atoi(inputStr)
		if err != nil {
			fmt.Printf("bookPrice Atoi error: %s", err)
			return
		}

		if bookPrice < 0 {
			fmt.Println("Invalid input. Please enter a non-negative integer.")
			return
		}

		totalPrice += bookPrice
	}

	var discountRate float64
	switch membershipLevel {
	case silverAcount:
		discountRate = silverDiscount
	case goldAccount:
		discountRate = goldDiscount
	case platinumAccount:
		discountRate = platinumDiscount
	}

	var reducedPrice float64
	reducedPrice = float64(totalPrice) * discountRate

	switch {
	case reducedPrice < 10:
		fmt.Printf("Add more items (%.2f) for free shipping\n", float64(totalPrice)-reducedPrice)
	case reducedPrice >= 10:
		println("You are eligible for free shipping")
	}

	fmtStr := `Summary for your order
	Orginal total price: %d
	Reduced totall price: %.2f`

	fmt.Printf(fmtStr, totalPrice, reducedPrice)

}
