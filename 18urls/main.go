package main

import (
	"fmt"
	"net/url"
)

var myUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Handling URLS in Golang")
	fmt.Println(myUrl)
	result, _ := url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println(qparams)
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])
	fmt.Println(qparams["paymentid"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "adi.dev",
		Path:    "/golang",
		RawPath: "user=adi",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
