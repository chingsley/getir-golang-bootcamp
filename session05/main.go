package main

import "fmt"

func main() {
	m1 := make(map[int]string)
	m2 := make(map[int]string, 10)
	m3 := map[int]string{1: "a" }
	fmt.Printf("m1 = %v\n", m1)
	fmt.Printf("m2 = %v\n", m2)
	fmt.Printf("m3 = %v\n", m3)

	// ***** Deleting a key-value pair *****
	m4 := map[int]string{1: "a"}
	fmt.Printf("%v\n", m4) // [1: a]
	delete(m4, 1)
	fmt.Printf("%v\n", m4) // []
	delete(m4, 2) // does nothing b/c m4 doesn't contain key 2
	fmt.Printf("%v\n", m4) // []

	// ***** getting lenth of map *****
	m5 := map[int]string{1: "a", 2: "b"}
	fmt.Printf("%d\n", len(m5)) // 2


	// ***** maps are passed by reference *****
	// if a map is modified in a function, applied changes will be visible on te mpa
	m6 := map[int]string{1: "a"}
	modifyMap(m6)
	fmt.Printf("%v\n", m6) // [2:b, 3:c]


	// ***** How to compare maps: see the mapsAreEqual function *****
	m7 := map[int]string{1: "a", 2: "b"}
	m8 := map[int]string{1: "a", 2: "b"}
	m9 := map[int]string{1: "a"}
	fmt.Println(mapsAreEqual(m7, m8)) // true
	fmt.Println(mapsAreEqual(m7, m9)) // false


	// ***** Maps with custom types *****
	type Person struct {
		Name string
		Age int
	}
	person1 := Person{Name: "Alice", Age: 30}
	person2 := Person{Name: "Bob", Age: 25}
	peopleMap := make(map[Person]string)
	peopleMap[person1] = "Friend"
	peopleMap[person2] = "Colleague"

	if person1 != person2 { // true becuase person1 is not equal to person2
		fmt.Println("person1 and person2 are not equal")
	}

	v := peopleMap[Person{Name: "Alice", Age: 30}]
	fmt.Printf("Alice is a %s", v) // Alice is a friend
}

func modifyMap(m map[int]string) {
	delete(m, 1)
	m[2] = "b"
	m[3] = "c"
}

func mapsAreEqual(m1, m2 map[int]string) bool {
	if len(m1) != len(m2) {
		return false
	}

	for key, val1 := range m1 {
		val2, ok := m2[key]
		if !ok || val1 != val2 {
			return false
		}
	}

	return true
}
