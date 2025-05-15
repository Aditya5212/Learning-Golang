package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())

}

func PerformPostJsonRequest() {
	url := "http://localhost:8000/post"
	// json payload
	requestBody := strings.NewReader(`
		{
			"Course Name" : "Mastery in Golang",
			"Price" : "100 percent efforts",
			"Platform" : "Youtube"
		}
	`)
	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"
	data := url.Values{}
	data.Add("firstname", "aditya")
	data.Add("lastname", "sharma")
	data.Add("email", "adi@gmail.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}

func main() {
	fmt.Println("Web Requests in Golang")
	// fmt.Println("Get Request")
	// PerformGetRequest()
	// fmt.Println("Post Request")
	// PerformPostJsonRequest()
	fmt.Println("Post Form Request")
	PerformPostFormRequest()
}
