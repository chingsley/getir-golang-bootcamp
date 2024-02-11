package main

import "fmt"

const (
	silverAccount = iota + 1
	goldAccount
	platinumAccount
	silverDiscount   = 0.9 // 10% discount for silver level
	goldDiscount     = 0.8 // 20% discount for gold level
	platinumDiscount = 0.7 // 30% discount for platinum level
)

func main() {
	bookTitle := "A Tale of Two Cities"
	var bookPrice int = 20

	var (
		membershipLevel int
		discountRate    float64
	)
	membershipLevel = 1

	if membershipLevel == silverAccount {
		discountRate = silverDiscount
	} else if membershipLevel == goldAccount {
		discountRate = goldDiscount
	} else if membershipLevel == platinumAccount {
		discountRate = platinumDiscount
	} else {
		fmt.Println("Invalid membership level.")
		return
	}

	var reducedPrice float64
	reducedPrice = float64(bookPrice) * discountRate

	fmtStr := `Summary of %s
	Original price: %d
	Reduced price: %.2f`

	fmt.Printf(fmtStr, bookTitle, bookPrice, reducedPrice)
}
