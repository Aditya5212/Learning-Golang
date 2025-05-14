package main

import "fmt"

func main() {
	fmt.Println("defer in Golang")

	// --- Basic Defer Example (Uncomment to see) ---
	// The 'defer' keyword schedules a function call (the "deferred" function)
	// to be executed just before the surrounding function (in this case, 'main') returns.
	//
	// If you uncomment these lines:
	//   defer fmt.Println("world") // This call is deferred.
	//   fmt.Println("hello")       // This call executes immediately.
	//
	// The output would be:
	//   defer in Golang
	//   hello
	//   world
	//
	// "world" is printed last because its call was deferred until 'main' is about to exit.

	// --- Defer in a Loop ---
	fmt.Println("counting")

	// When 'defer' is used inside a loop, the deferred function calls are pushed onto a stack.
	// They will be executed in Last-In, First-Out (LIFO) order after the surrounding
	// function ('main') finishes its other operations but before it actually returns.
	for i := 0; i < 10; i++ {
		// IMPORTANT: The arguments to a deferred function are evaluated when the 'defer'
		// statement itself is executed, NOT when the actual call happens.
		// So, in each iteration, 'fmt.Println(i)' is prepared with the *current* value of 'i'.
		// 1st iteration: defer fmt.Println(0) is scheduled.
		// 2nd iteration: defer fmt.Println(1) is scheduled (and pushed on top of the previous one).
		// ...and so on, up to defer fmt.Println(9).
		defer fmt.Println(i)
	}

	// This line will be printed BEFORE any of the deferred fmt.Println(i) calls.
	fmt.Println("done")
	// After "done" is printed, the deferred calls will execute in LIFO order:
	// 9
	// 8
	// 7
	// ...
	// 0
}
