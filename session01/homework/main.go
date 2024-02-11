package main

import (
	"fmt"
	"unicode/utf8"
)


const (
	Monday = iota + 1
	Tuesday
	Wednessday
	Thursday
	Friday
	Saturday
	Sunday
)

const MaxStringLength = 10

func main () {
	str1 := "red"
	str2 := "green"
	str3 := "blue"

	fmt.Println("--- Task 1: Variable Declaration and String Operations ---")
	fmt.Println("Length of str1: ", stringLength(str1))
	fmt.Println("Length of str2: ", stringLength(str2))
	fmt.Println("Length of str3: ", stringLength(str3))
	fmt.Println("")

	fmt.Println("--- Task 2: If-Else Conditions ---")
	compareStrings(str1, str2)
	compareStrings(str2, str3)
	compareStrings("john", "jane")
	fmt.Println("")


	fmt.Println("--- Task 3: Constants and IOTA ---")
	printDay("Monday")
	printDay("Tuesday")
	printDay("Wednessday")
	printDay("Thursday")
	printDay("Friday")
	printDay("Saturday")
	printDay("Sunday")


}

func stringLength(str string) int {
	return utf8.RuneCountInString(str)
}

func compareStrings(str1, str2 string) {
	if stringLength(str1) > stringLength(str2) {
		fmt.Printf("'%s' is longer than '%s'\n", str1, str2)
	} else if stringLength(str2) > stringLength(str1) {
		fmt.Printf("'%s' is longer than '%s'\n", str2, str1)
	} else {
		fmt.Printf("'%s' and '%s' have the same length\n", str1, str2)
	}
}

func printDay(day string) {
	if day == "Monday" {
		fmt.Println(day, ":", Monday)
	} else if day == "Tuesday" {
		fmt.Println(day, ":", Tuesday)
	} else if day == "Wednessday" {
		fmt.Println(day, ":", Wednessday)
	} else if day == "Thursday" {
		fmt.Println(day, ":", Thursday)
	} else if day == "Friday" {
		fmt.Println(day, ":", Friday)
	} else if day == "Saturday" {
		fmt.Println(day, ":", Saturday)
	} else if day == "Sunday" {
		fmt.Println(day, ":", Sunday)
	} else {
		fmt.Println("Invalid day of the week")
	}
}