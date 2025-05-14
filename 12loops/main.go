package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")

	// Declare and initialize a slice of strings named 'days'.
	// A slice is a dynamically-sized, flexible view into the elements of an array.
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// --- First type of for loop: C-style for loop ---
	// This is a classic for loop with three components separated by semicolons:
	// 1. init statement: 'i := 0' (executed before the first iteration)
	// 2. condition expression: 'i < len(days)' (evaluated before every iteration)
	// 3. post statement: 'i++' (executed at the end of every iteration)
	// 'len(days)' gives the number of elements in the 'days' slice.
	for i := 0; i < len(days); i++ {
		// 'days[i]' accesses the element at index 'i' in the 'days' slice.
		fmt.Println(days[i])
	}

	fmt.Println("This for loop is using Range")
	// --- Second type of for loop: for...range loop ---
	// The 'for...range' form iterates over a slice or map.
	// When ranging over a slice, two values are returned for each iteration:
	// 1. The index (here assigned to 'i').
	// 2. A copy of the element at that index (if you were to use it, e.g., 'for i, day := range days').
	// In this case, we are only interested in the index 'i' to access the element.
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for _, day := range days {
		fmt.Println(day)
	}

	rougeValue := 1
	for rougeValue < 10 {
		rougeValue++
		if rougeValue == 7 {
			goto lsi
		}
		if rougeValue == 5 {
			fmt.Println("Skipping 5")
			continue
		}
		if rougeValue == 8 {
			fmt.Println("Breaking the loop")
			break
		}
		fmt.Println(rougeValue)
	}

lsi:
	fmt.Println("Printing LSI Keywords at 7")

	// --- Third type of for loop: Go's "while" loop ---
	// Go doesn't have a dedicated 'while' keyword like some other languages.
	// Instead, a 'for' loop with only a condition is used to achieve the same behavior.
	// 'sum := 1' initializes a variable 'sum' to 1.
	sum := 1
	// This loop will continue as long as 'sum' is less than 10.
	for sum < 10 {
		// 'sum += sum' is shorthand for 'sum = sum + sum'. It doubles the value of 'sum'.
		sum += sum
		fmt.Println(sum)
	}

	valueMap := make(map[int]string)
	valueMap[1] = "One"
	valueMap[2] = "Two"
	valueMap[3] = "Three"
	valueMap[4] = "Four"
	valueMap[5] = "Five"

	for key, value := range valueMap {
		fmt.Printf("Key is %v and value is %q\n", key, value)
	}
	// for key := range valueMap {
	// 	fmt.Printf("Key is %v and value is %q\n", key, valueMap[key])
	// }
}
