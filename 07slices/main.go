package main

// Importing necessary packages:
// "fmt" for formatted input/output operations (like printing).
// "sort" for sorting collections like slices.
import (
	"fmt"
	"sort"
)

// The main function, where program execution begins.
func main() {
	fmt.Println("Slices in Golang")

	// Declare and initialize a slice of strings.
	// Unlike arrays, slices don't have a fixed size defined at compile time.
	// This is a common way to create a slice with initial values.
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	// Print the type of fruitList. This will show "[]string", indicating it's a slice of strings.
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	// Append new elements to the fruitList slice.
	// The 'append' function returns a new slice containing the original elements plus the new ones.
	// It's important to assign the result of 'append' back to the original slice variable.
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	// Slice the fruitList. This creates a new slice that references a portion of the original slice's underlying array.
	// fruitList[1:3] takes elements from index 1 up to (but not including) index 3.
	// So, it will take "Tomato" and "Peach".
	// The original fruitList is then reassigned to this new, smaller slice.

	// fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	// Create a slice of integers using the 'make' function.
	// 'make([]int, 4)' creates a slice of integers with an initial length of 4.
	// Its elements are initialized to their zero value (0 for integers).
	highScores := make([]int, 4)

	// Assign values to the elements of the highScores slice using their indices.
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	fmt.Println(highScores)

	// The following line would cause a runtime panic: index out of range.
	// Slices have a length, and you can't access an index equal to or greater than the length.
	// highScores[4] = 777

	// You can append to a slice made with 'make' just like any other slice.
	// highScores = append(highScores, 555, 666, 321)

	// Sort the highScores slice in ascending order.
	// The 'sort.Ints' function modifies the slice in place.
	sort.Ints(highScores)
	fmt.Println(highScores)
	// Check if the highScores slice is sorted and print the boolean result.
	fmt.Println(sort.IntsAreSorted(highScores))

	// // How to Remove a value from slice by Index (This entire block is currently commented out)

	// Declare and initialize a slice of strings named 'courses'.
	var courses = []string{"0reactjs", "1javascript", "2swift", "3python", "4ruby"}
	// Print the initial 'courses' slice.
	fmt.Println(courses)
	// Define an integer variable 'index' and set its value to 2.
	// This is the index of the element we intend to remove ("2swift").
	var index int = 2
	// Remove the element at the specified 'index'.
	// This is done by creating two new slices: one from the beginning up to 'index' (courses[:index]),
	// and another from 'index + 1' to the end (courses[index+1:]).
	// These two slices are then appended together, effectively skipping the element at 'index'.
	// The '...' (ellipsis) is used to unpack the elements of the second slice (courses[index+1:])
	// so they can be appended individually.
	courses = append(courses[:index], courses[index+1:]...)
	// Print the 'courses' slice after removing the element.
	fmt.Println(courses)
}
