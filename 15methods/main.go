package main

import "fmt"

type User struct {
	name     string // Field 'name' of type string, to store the user's name.
	email    string // Field 'email' of type string, to store the user's email address.
	loggedIn bool   // Field 'loggedIn' of type bool, to indicate if the user is currently logged in.
	age      int    // Field 'age' of type int, to store the user's age.
}

func (u User) GetStatus() {
	fmt.Println("Is user logged in: ", u.loggedIn)
}

func (u User) NewEmail() {
	u.email = "Aditya@gmail.com"
	fmt.Println("New Email of this user is: ", u.email)
}

func main() {
	fmt.Println("Methodes in Golang")

	aditya := User{
		name:     "aditya",                   // Assigning "aditya" to the 'name' field.
		email:    "adityadange135@gmail.com", // Assigning an email address to the 'email' field.
		age:      20,                         // Assigning 20 to the 'age' field.
		loggedIn: true,                       // Assigning true to the 'loggedIn' field.
		// A comma is required after the last field if the closing brace '}' is on a new line.
	}
	// Print the details of the 'aditya' struct.
	// '%+v' is a format specifier that prints the struct with field names.
	fmt.Println("Aditya Details are: ", aditya)
	fmt.Printf("Aditya Details are: %+v\n", aditya)
	fmt.Printf(" name: %v\n email: %v\n age: %v\n loggedIn: %v\n", aditya.name, aditya.email, aditya.age, aditya.loggedIn)

	// To find the type of a variable, you can use fmt.Printf with the %T verb.
	// '%T' is a format specifier that prints the type of the variable.
	fmt.Printf("The type of aditya is: %T\n", aditya)

	aditya.GetStatus()
	aditya.NewEmail()
	fmt.Println(aditya.email)
}
