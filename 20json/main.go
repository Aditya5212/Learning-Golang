package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string
	Price    float32
	Platform string
	Password string
	Tags     []string
}

func main() {
	fmt.Println("Json in Golang")
	EasyCourses := []Course{
		{"ReactJS", 299, "Aditya's Mastercalss", "abc123", []string{"web-dev", "js"}},
		{"MERN", 1299, "Aditya's Mastercalss", "bcd123", []string{"web-dev", "js"}},
		{"Angular", 1299, "Aditya's Mastercalss", "def123", []string{"web-dev", "js"}},
	}

	fmt.Printf("%#v\n", EasyCourses)

	// package this data as Json - Marshaling
	fileJson, err := json.Marshal(EasyCourses)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n\n", fileJson)
	fmt.Printf("%T\n\n", fileJson[0])
	// fmt.Println("fileJson[1] : ", string(fileJson[0:]))

}
