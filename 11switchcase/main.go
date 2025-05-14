package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Prints a message to the console.
	fmt.Println("Switch and case in Golag")

	// Seeding the random number generator.
	// 'rand.Seed' initializes the default Source to a deterministic state.
	// If Seed is not called, the generator behaves as if seeded by Seed(1).
	// 'time.Now().UnixNano()' provides a seed value that changes very frequently (nanoseconds since epoch),
	// which helps in generating different sequences of random numbers each time the program runs.
	// Without a unique seed, 'rand.Intn' would produce the same sequence of numbers every run.
	rand.Seed(time.Now().UnixNano())

	// Generating a random integer.
	// 'rand.Intn(6)' generates a random integer in the range [0, 6-1], so [0, 5].
	// We add 1 to make it a typical dice roll value, i.e., in the range [1, 6].
	// The ':= ' syntax is a short variable declaration, used when declaring a variable inside a function.
	// It infers the type of 'diceNumber' (which will be 'int') from the value on the right.
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	// The 'switch' statement provides a multi-way branch.
	// It evaluates an expression ('diceNumber' in this case) and then
	// compares it with the values in each 'case' clause.
	switch diceNumber {
	case 1: // If diceNumber is 1
		fmt.Println("Dice value is 1 and you can open")
	case 2: // If diceNumber is 2
		fmt.Println("You can move 2 spot")
	case 3: // If diceNumber is 3
		fmt.Println("You can move to 3 spot")
		fallthrough // The 'fallthrough' statement transfers control to the next case clause.
		// In Go, cases do not automatically fall through like in C or Java.
		// You must explicitly use 'fallthrough' if you want that behavior.
		// So, if diceNumber is 3, it will print "You can move to 3 spot" AND then also execute case 4.
	case 4: // If diceNumber is 4, or if it was 3 and fell through
		fmt.Println("You can move to 4 spot")
	case 5: // If diceNumber is 5
		fmt.Println("You can move to 5 spot")
	case 6: // If diceNumber is 6
		fmt.Println("You can move to 6 spot and roll dice again")
	default: // The 'default' case is executed if none of the other cases match.
		// In this specific scenario, since diceNumber is always between 1 and 6,
		// this default case should ideally not be reached.
		// However, it's good practice to include a default case for unexpected values.
		fmt.Println("What was that!")
	}
}
