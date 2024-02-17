package main

import (
	"fmt"
)

func main() {
	age := 25

	switch {
	case age < 18:
		fmt.Println("You are a minor")
	case age >= 18 && age < 60:
		fmt.Println("You are an adult.")
	case age >= 60:
		fmt.Println("You are a senior citizen.")
	default:
		fmt.Println("Invalid age.")
	}
}
