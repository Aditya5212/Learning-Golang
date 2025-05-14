package main

import "fmt"

func content() {
	fmt.Println("---------------------> Functions in Golang <---------------------")
}

func adder(a int, b int) (sum int) {
	sum = a + b
	return
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func main() {
	content()
	firstVal, secondVal := 0, 0
	fmt.Println("Enter First numbers: ")
	// fmt.Scanln() returns two values: the number of items scanned and an error.
	fmt.Scanln(&firstVal)
	fmt.Println("Enter First numbers: ")
	fmt.Scanln(&secondVal)
	fmt.Println("The sum is: ")
	result := adder(firstVal, secondVal)
	println(result)

	proResult := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("The proResult is: ", proResult)
}
