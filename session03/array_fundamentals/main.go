package main

import "fmt"

func main() {
	x := [5]int{10, 20, 30, 40, 50} // initialized with values
	var y [5]int = [5]int{10, 20, 30} // partial assignment

	fmt.Println(x) // [10, 20, 30, 40, 50]
	fmt.Println(y) // [10, 20, 30, 0, 0]

	//****************** 2-D array ***************
	sample := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	// traversing 2-D array:
	for _, row := range sample {
		for _, val := range row {
			fmt.Println(val)
		}
	}

	//****************** Keyed Elements ***************
	weekend := []bool{5: true, 6: true} // the rest will be false
	isWeekend := []int{5: 1, 6: 1} // the rest will be 0
	fmt.Println(weekend, isWeekend)


	//****************** Unnanmed (Composite, Unassigned Arrays ***************
	type Beyza [3]int
	arr1 := Beyza{1, 2, 3}
	//fmt.Println(arr1)

	arr2 := [3]int{1, 2, 3}

	fmt.Println(arr1 == arr2) // true

	type Swat [3]int
	arr3 := Swat{1, 2, 3}
	//fmt.Println(arr1 == arr3) // will cause an error because they're of different types
	
	fmt.Println(arr1 == Beyza(arr3))// true

	numbers := [5]int{10, 20, 30, 40}
	fmt.Println(numbers) // [10, 20, 30, 40, 0], fifth element is zero

	copyOfNumbers := numbers // arrays are copied by value (numbers != copyOfNumbers)
	fmt.Println(copyOfNumbers == numbers) // true

	copyOfNumbers[0] = 100 // change first element in copy
	fmt.Println(copyOfNumbers == numbers) // false
	fmt.Println(numbers[0], copyOfNumbers[0]) // 10, 100


	// ** MATRIX **//
	matrix := [3][3]string {
		{"A-A", "A-B", "A-C"},
		{"B-A", "B-B", "B-C"},
		{"C-A", "C-B", "C-C"},
	}

	for i, row := range matrix {
		for j, cell := range row {
			fmt.Println("Row: ", i, ", Column: ",j, ", Cell:", cell)
		}
	}


	// ** KEYED ARRAY **//
	names := [3]string{"Alice", "Bob", "Charlie"}
	keyedNames := [...]string{2: "Charlie", 0: "Alice", 1: "Bob"}
	fmt.Println(names == keyedNames) // true

}
