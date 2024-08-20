package main

import "fmt"

func main() {
	// ******** Copying slices *************
	a := []int{10, 20, 30, 40}
	b := []int{1, 1, 1, 1, 1}
	fmt.Printf("a: %v, b: %v\n", a, b)

	copy(b, a) // copy a into b, ie. set b = a
	fmt.Printf("a: %v, b: %v\n", a, b) // b = [10, 20, 30, 40, 1] Copies the four elements of a and includes the last 1 from b to maintain a length of 5

	fmt.Println("----------------")
	a = []int{10, 20, 30, 40}
	b = []int{1, 1}
	fmt.Printf("a: %v, b: %v\n", a, b)

	copy(b, a) // copy a into b
	fmt.Printf("a: %v, b: %v\n", a, b) // b = [10, 20], maintains the length of 2

	fmt.Println("----------------")
	a = []int{10, 20, 30, 40}
	b = make([]int, 4)
	copy(b, a)
	fmt.Println("b: ", b) // [10, 20, 30, 40]



	fmt.Println("----------------")
	// ******** Length and Capacity *************
	s1 := []int{1, 2, 3}
	fmt.Printf("s1: %v, length: %d, capacity: %d\n", s1, len(s1), cap(s1)) // length: 3, capacity: 3

	s2 := make([]int, 3)
	fmt.Printf("s2: %v, length: %d, capacity: %d\n", s2, len(s2), cap(s2)) // length: 3, capacity: 3

	S3 := make([]int, 0, 3)
	fmt.Printf("S3: %v, length: %d, capacity: %d\n", S3, len(S3), cap(S3)) // length: 0, capacity: 3


	s4 := make([]int, 2, 3)
	fmt.Printf("s4: %v, length: %d, capacity: %d\n", s4, len(s4), cap(s4)) // [0, 0], length: 3, capacity: 3
	s4 = append(s4, 4)
	fmt.Printf("s4: %v, length: %d, capacity: %d\n", s4, len(s4), cap(s4)) // [0, 0, 4], length: 3, capacity: 3
	s4 = append(s4, 5) // slice will be resized
	fmt.Printf("s4: %v, length: %d, capacity: %d\n", s4, len(s4), cap(s4)) // [0, 0, 4, 5], length: 4, capacity: 6
	s4 = append(s4, 6, 7, 8) // slice will be resized
	fmt.Printf("s4: %v, length: %d, capacity: %d\n", s4, len(s4), cap(s4)) // [0, 0, 4, 5, 6, 7, 8], length: 7, capacity: 12
}
