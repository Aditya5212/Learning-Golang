package main

import "fmt"

func changeValue(str *string) {
	*str = "changed"
}

func changeValue2(str string) {
	str = "changed"
}

// & *
func main() {
	// toChange := "test"
	// changeValue2(toChange)
	// println(toChange)
	// changeValue(&toChange)
	// println(toChange)

	// x := 7
	// y := &x
	// fmt.Println(&x)
	// fmt.Println(y)
	// fmt.Printf("%T\n", y)
	// fmt.Println(*y)
	// *y = 10 // Derefrence
	// fmt.Println(x)
	// fmt.Println(*y)

	str := "testString"
	pointer := &str
	fmt.Println("Pointer: ", pointer)
	fmt.Println("Value: ", *pointer)
	fmt.Println("Pointer Reference: ", &pointer)
	fmt.Println("Value Reference *&pointer: ", *&pointer)
	fmt.Println("Value Reference &*pointer: ", &*pointer)
	if *&pointer == pointer {
		fmt.Println("True")
	}

}
