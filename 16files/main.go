package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text data inside the file is \n", string(databyte))
}

func main() {
	fmt.Println("File Handling in GoLang")
	content := "This needs to go in a file - File Data"
	file, err := os.Create("./mylocalfile.txt")
	checkNilErr(err)
	length, err := file.WriteString(content)
	checkNilErr(err)
	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./mylocalfile.txt")

}
