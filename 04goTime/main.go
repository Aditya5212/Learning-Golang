package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of Golang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006"))
	createDate := time.Date(2020, time.November, 10, 24, 0, 0, 0, time.UTC)
	fmt.Println(createDate)
	numcpu := runtime.NumCPU()
	fmt.Println(numcpu)
}
