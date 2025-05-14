package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "https://catfact.ninja/fact"

func main() {
	fmt.Println("Web Requests in Golang")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type: %T\n", response)
	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)
	fmt.Println("Content Type: ", response.Header.Get("Content-Type"))
	// fmt.Println("Response Headers: ", response.Header)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)
	fmt.Println("Response String: ", responseString.String())
	fmt.Println("Byte Count: ", byteCount)
	defer response.Body.Close()
}
