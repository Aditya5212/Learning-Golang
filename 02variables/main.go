package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Public var init
var IsLogin bool = true

// add function: Demonstrates named return values and a naked return.
// In Go, function return types can be named. Here, the return type is 'int' and is named 'sum'.
// A return statement without any arguments returns the current value of the named return variable.
func add(x int, y int) (sum int) {
	sum = x + y // This is called a "naked return" because we're relying on the named return value.
	return
}

func main() {
	// Declaring and initializing multiple variables of the same type in a single line.
	var c, python, java bool = true, false, true
	// Using short variable declaration (:=) to declare and initialize local variables.
	// Go infers the types based on the assigned values.
	// No Var style
	golang, javascript := true, false
	fmt.Println(golang, javascript)

	// Declaring a string variable.
	var username string = "Aditya Dange"
	fmt.Println("Hello", username)

	// Generating and printing a random integer between 0 (inclusive) and 10 (exclusive).
	fmt.Println("This is my favorite number ", rand.Intn(10))
	// Calculating and printing the square root of 16.
	fmt.Println("Square Root of 16 is ", math.Sqrt(16))
	// Accessing and printing the constant Pi from the math package.
	// Note: Identifiers (functions, constants, types, variables) exported from a package (i.e., accessible outside the package) must start with a capital letter.
	fmt.Println(math.Pi)
	fmt.Println(add(42, 58))
	fmt.Println(c, python, java)

	fmt.Println("\nVariables declared without a explicit return value are given their zero value")
	var i int     // Integer variables default to 0
	var f float64 // Float variables default to 0.0
	var b bool    // Boolean variables default to false
	var s string  // String variables default to "" (empty string)
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	fmt.Println("\nType Conversion")
	var x, y int = 3, 4
	var f1 float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f1)
	fmt.Println(x, y, z)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Println("\nConstants") // Constants can not be declared using := syntar
	// Constants are declared using the `const` keyword.
	// They represent values that are known at compile time and cannot be changed during program execution.
	// Unlike variables, constants cannot be declared using the short variable declaration (:=) syntax.
	const World = "World"
	fmt.Println("Hello", World)

}
