package main

import "fmt"

// IShape is an interface with two definitions
type IShape interface {
	Area() float64
	Perimeter() float64
	IWhat // a interface contains other interface, so the implementer has to implement all the methods
}

type IWhat interface {
	What() string
}

func Shape(ishape IShape) { // dependency injection
	fmt.Printf("Area of %v: %.3f\n", ishape.What(), ishape.Area())
	fmt.Printf("Perimeter of %v: %.3f\n", ishape.What(), ishape.Perimeter())
	fmt.Println()
}
