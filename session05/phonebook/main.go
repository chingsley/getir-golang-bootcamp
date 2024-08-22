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

const MaxEntries = 100

var phonebook = make(map[string]string, MaxEntries)

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

		switch choice {
		case 1:
			addContact()
		case 2:
			deleteContact()
		case 3:
			searchContact()
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

func addContact() {
	if len(phonebook) >= MaxEntries {
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

	if _, exists := phonebook[name]; exists {
		fmt.Printf("Conttact with name %s already exists\n", name)
		return
	}

	fmt.Print("Enter a number: ")
	_, errNumber := fmt.Scan(&number)
	if errNumber != nil {
		fmt.Println(errNumber.Error())
		return
	}

	phonebook[name] =  number
	fmt.Println("contact added successfully")
}

func deleteContact() {
 var name string

 fmt.Print("Enter a name: ")
 _, err := fmt.Scan(&name)
 if err != nil {
 	fmt.Println(err.Error())
 	return
 }

 if _, exists := phonebook[name]; exists {
 	delete(phonebook, name)
 	fmt.Println("Contact deleted!")
 } else {
	 fmt.Println("contact not found")
 }
}

func searchContact() {
	var name string

	fmt.Print("Enter a name: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if _, exists := phonebook[name]; exists {
		fmt.Printf("Contact found: %s -> %s", name, phonebook[name])
	} else {
		fmt.Println("contact not found")
	}
}

func listPhoneBook() {
 for name, number := range phonebook {
 	fmt.Println(name, " -> ", number)
 }
}
