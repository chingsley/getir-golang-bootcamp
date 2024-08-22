package main

/*
 THE PHONEBOOK APPLICATION:
- get the options from cli
- add a contact
- delete a contact
- search a contact
- list the phonebook elements
 */

import "fmt"

const MaxEntries = 4

var phonebook [MaxEntries][2]string
var currentIndex = 0

func main() {
	for {
		fmt.Println("\nPhonebook Menu:")
		fmt.Println("1. Add new contact")
		fmt.Println("2. Delete a new contact")
		fmt.Println("3. Search for a contact")
		fmt.Println("4. List the phonebook")
		fmt.Println("5. Exit")

		var choice int

		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(phonebook[0], phonebook[0][0], phonebook[0][1], phonebook[0][1] == "")

		switch choice {
		case 1:
			addContact(&currentIndex)
		case 2:
			deleteContact(&currentIndex)
		case 3:
			searchContact(currentIndex)
		case 4:
			listPhoneBook()
		case 5:
			fmt.Println("Exiting the phonebook")
			return
		default:
			fmt.Println("Invalid option selected, please choose a number from 1 to 5")
		}
	}
}

func addContact(currIdx *int) {
	if *currIdx >= MaxEntries {
		fmt.Println("Phonebook is full. Cannot add new contact")
		return
	}
	var name, number string

	fmt.Print("Enter a name: ")
	_, errName := fmt.Scan(&name)
	if errName != nil {
		fmt.Println(errName.Error())
		return
	}

	fmt.Print("Enter a number: ")
	_, errNumber := fmt.Scan(&number)
	if errNumber != nil {
		fmt.Println(errNumber.Error())
		return
	}

	phonebook[*currIdx][0] =  name
	phonebook[*currIdx][1] = number
	fmt.Println("contact added successfully")
	*currIdx++
}

func deleteContact(currentIdx *int) {
 var name string

 fmt.Print("Enter a name: ")
 _, err := fmt.Scan(&name)
 if err != nil {
 	fmt.Println(err.Error())
 	return
 }

 for i := 0; i < *currentIdx; i++ {
 	if phonebook[i][0] == name {
 		phonebook[i][0] = ""
 		phonebook[i][1] = ""
 		fmt.Println("Contact deleted successfully")
 		*currentIdx--
 		return
	}
 }

 fmt.Println("contact not found")
}

func searchContact(currentIdx int) {
	var name string

	fmt.Print("Enter a name: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for i := 0; i < currentIdx; i++ {
		if phonebook[i][0] == name {
			fmt.Printf("Contact found: %s -> %s", phonebook[i][0], phonebook[i][1])
			return
		}
	}

	fmt.Println("contact not found")
}

func listPhoneBook() {
 for _, contact := range phonebook {
 	fmt.Println(contact[0], " -> ", contact[1])
 }
}
