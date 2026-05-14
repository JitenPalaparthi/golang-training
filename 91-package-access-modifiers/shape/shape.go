package shape

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

type circle struct { // c starts with lowercase
	R    float32
	a, p float64
}

func Newcirlce(r float32) *circle { // Newcircle is Uppercase
	return &circle{r, 0, 0}
}

type Circle struct { // c starts with Uppercase
	R float32 // R is with Uppercase
	a float64 // a is restricted
	P float64 // P is unrestricted
}

func greet() { // restricted to outside
	println("Welcome to shape package")
}

func Greet() { // unrestricted can be called
	greet() // greet can be called locally but not out of this package
}
