package main

import "fmt"

type Point struct {
	X int
	Y int
}

type Circle struct {
	center *Point
	Radius int
}

func changeX(pt *Point, val int) {
	pt.X = val
}

func main() {
	pt1 := &Point{1, 2}
	circle := Circle{pt1, 7}
	// var pt2 Point = Point{Y: 3, X: 4}
	changeX(pt1, 10)
	fmt.Println(circle.center.X)
	fmt.Println(circle.center.Y)
	fmt.Println(circle.Radius)

}
