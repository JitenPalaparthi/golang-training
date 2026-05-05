package main

import "fmt"

// IShape is an interface with two definitions
type IShape interface {
	Area() float64
	Perimeter() float64
}

func Shape(ishape IShape) { // dependency injection
	fmt.Printf("Area:%.3f\n", ishape.Area())
	fmt.Printf("Perimeter:%.3f\n", ishape.Perimeter())
	fmt.Println()
}
