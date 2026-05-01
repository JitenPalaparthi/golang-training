package main

import "fmt"

func main() {

	r1 := NewRect(10.4, 12.5)
	a1 := r1.Area()
	p1 := r1.Perimeter()

	fmt.Printf("Area of Rect:%.2f\nPerimeter of Rect:%.2f\n", a1, p1)
	fmt.Printf("Area of Rect:%.2f\nPerimeter of Rect:%.2f\n", r1.A, r1.P)

	s1 := NewSquare(23.4)

	a2, p2 := s1.Area(), s1.Perimeter()

	fmt.Printf("Area of Square:%.2f\nPerimeter of Square:%.2f\n", a2, p2)

}

// How to create a constructor ? Go does not follow typical OOPS concepts
