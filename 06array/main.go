// Every Go program starts with a package declaration.
// The 'main' package is special: it tells the Go compiler that this package should compile into an executable program.
package main

// 'fmt' (format) is a standard Go package for formatted I/O (like printing to the console).
import "fmt"

// The 'main' function is the entry point of a 'main' package executable.
func main() {
	fmt.Println("Array in Golang")

	// Declare an array named 'fruitList'.
	// In Go, an array has a fixed size, which is part of its type.
	// This array can hold 4 elements, and each element must be a 'string'.
	// When an array is declared but not initialized, its elements are set to their zero value.
	// For strings, the zero value is an empty string "".
	var fruitList [4]string

	// Assign values to specific elements of the array using their index.
	// Array indices in Go are 0-based.
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	// fruitList[2] is not assigned, so it will remain its zero value (empty string).
	fruitList[3] = "Peach"

	// Print the entire 'fruitList' array.
	fmt.Println("Fruit list is: ", fruitList)
	// Print the length of the 'fruitList' array. The 'len()' function returns the number of elements.
	fmt.Println("Fruit list is: ", len(fruitList))

	// Declare and initialize an array named 'vegList' in one step using an array literal.
	// The size [3] is explicitly stated.
	// The array will contain "Potato", "Beans", and "Mushroom".
	var vegList = [3]string{"Potato", "Beans", "Mushroom"}
	fmt.Println("Vegy list is: ", vegList)
	fmt.Println("Vegy list is: ", len(vegList))
}
