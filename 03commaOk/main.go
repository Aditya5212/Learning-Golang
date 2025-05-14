package main

import (
	// `bufio` package provides buffered I/O operations.
	// We'll use it here for reading user input line by line.
	"bufio"
	// `fmt` package implements formatted I/O, similar to C's printf and scanf.
	"fmt"
	// `os` package provides a platform-independent interface to operating system functionality.
	// We'll use it for accessing standard input (os.Stdin).
	"os"
	// `strconv` package implements conversions to and from string representations of basic data types.
	"strconv"
	// `strings` package implements simple functions to manipulate UTF-8 encoded strings.
	"strings"
)

// The `main` function is the entry point of a Go executable program.
func main() {
	// Declare and initialize a string variable `welcome`.
	// `:=` is the short variable declaration operator. It infers the type of `welcome` (string).
	welcome := "Welcome to user input!"
	// Print the welcome message to the console.
	fmt.Println(welcome)

	// Create a new buffered reader that reads from standard input (os.Stdin).
	// This allows us to read text entered by the user in the console.
	reader := bufio.NewReader(os.Stdin)

	// --- Comma OK syntax ---
	// `reader.ReadString('\n')` reads input until the first occurrence of '\n' (newline character).
	// It returns two values: the string read (including the newline) and an error.
	// Here, we are using the "comma ok" or "comma error" idiom.
	// `input` will store the string, and `_` (the blank identifier) is used to discard the error value.
	// We are choosing to ignore the potential error from ReadString in this specific case for simplicity.
	// NOTE: When using the comma-ok or comma-error syntax like `value, err := ...`,
	// the `:=` operator is crucial. It declares new variables on the left-hand side.
	// If `err` was already declared in the scope, and `input` is new, `:=` still works.
	// If both `input` and `err` were already declared, you would use `=` instead.
	input, _ := reader.ReadString('\n')
	fmt.Println("This is user provided input :", input)

	fmt.Println("Rate me out of 5")
	// Read the rating input from the user, similar to the previous input.
	// Again, we're using the blank identifier `_` to ignore the error from ReadString.
	rating, _ := reader.ReadString('\n')
	fmt.Println("Your rating is :", rating)

	// --- Converting String to Number and Error Handling ---
	// `strings.TrimSpace(rating)` removes leading and trailing white space,
	// including the newline character `\n` that `ReadString` includes.
	// `strconv.ParseFloat` attempts to convert the cleaned string `rating` into a float64 number.
	// This is another example of the "comma error" idiom.
	// `numRating` will hold the converted float if successful, and `err` will hold any error that occurred.
	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	// Check if an error occurred during the conversion.
	if err != nil {
		// If `err` is not `nil`, an error happened (e.g., user entered "hello" instead of a number).
		fmt.Println(err)
	} else {
		// If `err` is `nil`, the conversion was successful.
		// We can now use `numRating` as a float.
		fmt.Println("Added 1 to your rating :", numRating+1)
	}
}
