package main

import "fmt"

func main() {
	fmt.Println("Pointers in Golang")

	// pointers - Memory Address
	// Declaring a pointer variable 'ptr' that can store the memory address of an integer.
	// By default, a pointer that doesn't point to anything has a 'nil' value.
	var ptr *int
	fmt.Println(ptr) //-> <nil>

	// Declare an integer variable 'myNum' and initialize it with 23.
	myNum := 23

	// Declare a pointer 'ptr2' and store the memory address of 'myNum' in it.
	// The '&' operator is used to get the memory address of a variable.
	var ptr2 = &myNum
	fmt.Println("Value of actual pointer is", ptr2)  // This will print the memory address.
	fmt.Println("Value of actual pointer is", *ptr2) // This will print the value stored at that memory address (dereferencing).

	// Dereference 'ptr2' to get the value it points to (which is myNum's value).
	// Multiply that value by 2.
	// Store the result back into the memory location pointed to by 'ptr2'.
	// This directly modifies the original 'myNum' variable.
	*ptr2 = *ptr2 * 2
	fmt.Println("New value is: ", myNum) //->46, because myNum was modified through the pointer.
}
