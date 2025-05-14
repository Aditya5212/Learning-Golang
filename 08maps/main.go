package main

import (
	"fmt"
)

// isRust checks if "Rust" (with key "RS") exists in the provided map.
// If it doesn't exist, it adds "Rust" to the map and returns false.
// If it already exists, it prints a message and returns true.
// The function modifies the map directly because maps are reference types in Go.
func isRust(languages map[string]string) bool {
	// This is the "comma ok" idiom to check if a key exists in a map.
	// 'languages["RS"]' attempts to retrieve the value associated with the key "RS".
	// If the key "RS" exists, 'ok' will be true, and '_' (the blank identifier) would hold the value.
	// If the key "RS" does not exist, 'ok' will be false, and '_' would be the zero value for the map's value type (an empty string "" in this case).
	if _, ok := languages["RS"]; !ok { // Check if "RS" does NOT exist
		// If "RS" is not in the map, add it with the value "Rust".
		languages["RS"] = "Rust"
		fmt.Println("Rust added")
		// Return false because Rust was not initially present.
		return false
	} else {
		// If "RS" is already in the map (i.e., ok is true).
		fmt.Println("Rust already exists")
		// Return true because Rust was already present.
		return true
	}
}

// The 'main' function is the entry point of the executable program.
func main() {
	fmt.Println("Maps in Golang")

	// Declare and initialize an empty map called 'languages'.
	// 'make(map[string]string)' creates a map where keys are strings and values are strings.
	languages := make(map[string]string)

	// Add key-value pairs to the 'languages' map.
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["GO"] = "Golang"

	fmt.Println("List of All Languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	// Delete a key-value pair from the map using the 'delete' built-in function.
	delete(languages, "RB")
	fmt.Println("List of All Languages: ", languages)

	// Call the isRust function, passing the 'languages' map.
	// The 'isRust' function might modify the 'languages' map.
	result := isRust(languages)
	fmt.Println("Rust exists: ", result)
	fmt.Println("List of All Languages: ", languages)

	// To check if a key exists, use the "comma ok" idiom.
	// val will be the value if the key exists, or the zero value (empty string for string) if not.
	// ok will be true if the key exists, and false otherwise.

	// Loop through the map using a 'for...range' loop.
	// In each iteration, 'key' gets the current key and 'value' gets the corresponding value.
	// The order of iteration over a map is not guaranteed in Go.
	for key, value := range languages {
		fmt.Printf("Key: %v Value: %v\n", key, value)
	}
}
