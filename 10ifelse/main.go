package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println("If Else in Golang")

	// Declare an integer variable 'loginCount'. It will be initialized to its zero value, which is 0 for int.
	var loginCount int

	// Prompt the user to enter a number.
	fmt.Println("Enter the number of logins:")
	// 'fmt.Scan' reads formatted input from standard input (the console)
	// and stores it in the variable pointed to by '&loginCount'.
	// 'fmt.Scan' returns the number of items successfully scanned and an error.
	// Note: The error from fmt.Scan itself is not being checked here.
	fmt.Scan(&loginCount)

	// This line prints the value of loginCount to the console.
	// 'fmt.Println' returns the number of bytes written and an error.
	// We are capturing the error in 'err' and discarding the number of bytes written using '_'.
	_, err := fmt.Println("You entered:", loginCount) // Added "You entered:" for clarity

	// Check if an error occurred while trying to print 'loginCount'.
	// This is generally not how you'd check for errors from 'fmt.Scan'.
	// An error here would typically mean an issue with standard output, not the input process itself.
	if err != nil {
		fmt.Println("Error printing the login count:", err)
	} else {
		fmt.Println("Successfully printed the login count.")
	}

	// Declare a string variable 'result' to store a message based on 'loginCount'.
	var result string

	// This is an 'if-else if-else' statement.
	// It checks conditions sequentially and executes the block associated with the first true condition.
	if loginCount < 10 {
		result = "Regular User" // If loginCount is less than 10.
	} else if loginCount > 10 {
		result = "Watch Out" // If loginCount is greater than 10 (and not less than 10).
	} else {
		result = "Exactly 10 login count" // If loginCount is neither less than nor greater than 10 (i.e., it's 10).
	}
	fmt.Println(result)

	// This is a simple 'if-else' statement.
	// It checks if 'loginCount' is even or odd using the modulo operator '%'.
	if loginCount%2 == 0 {
		fmt.Println("Number is even") // If the remainder of loginCount divided by 2 is 0.
	} else {
		fmt.Println("Number is odd") // Otherwise.
	}

	// This 'if' statement demonstrates a short statement.
	// 'num := 3' declares and initializes 'num' within the scope of this 'if-else' block.
	// The condition 'num < 10' is then evaluated.
	if num := 3; num < 10 {
		fmt.Println("num (which is 3) is Small")
	} else {
		fmt.Println("num (which is 3) is Large") // This 'else' block won't be reached with num = 3.
	}

	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
