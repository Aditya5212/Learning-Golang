package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"math"
)

type Rectangle struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Shape interface {
	Area() float64
}

func Area(s Shape) float64 {
	return s.Area()
}

func main() {
	fmt.Println("Interfaces in Golang")
	c1 := Circle{radius: 5.5}
	r1 := Rectangle{length: 10, width: 5}
	shapes := []Shape{c1, r1}
	for _, shape := range shapes {
		fmt.Println(shape.Area())
	}
	fmt.Println("Area of Circle:", c1.Area())
	fmt.Println("Area of Rectangle:", r1.Area())
	// fmt.Println("Area of Circle Area(c1):", Area(c1))
	// fmt.Println("Area of Rectangle Area(r1):", Area(r1))

	// payload := []byte("Hello High Value Software Engineer")
	// hashAndBroadcast(bytes.NewReader(payload))
}

type HashReader struct {
	bytes.Reader
	buf *bytes.Buffer
}

func NewHashReader(b []byte) *HashReader {
	return &HashReader{
		Reader: *bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}

func (h *HashReader) hash() string {
	return hex.EncodeToString(h.buf.Bytes())
}

func hashAndBroadcast(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	hash := sha1.Sum(b)
	// fmt.Println(hash)
	fmt.Println(hex.EncodeToString(hash[:]))

	return broadcast(r)
}

func broadcast(r io.Reader) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	fmt.Println("String of Bytes", string(b))
	return nil

}
