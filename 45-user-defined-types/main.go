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

type Rect struct {
	B, H float32
	A, P float64
}

type Square struct {
	S float32
}

// new is a function can think it as a constructor
func NewRect(b, h float32) *Rect {
	return &Rect{B: b, H: h}
}

func NewSquare(s float32) *Square {
	return &Square{s}
}

func (r *Rect) Area() float64 { // pass by ref or call by ref
	r.A = float64(r.B) * float64(r.H)
	return r.A
}

func (r *Rect) Perimeter() float64 { // pass by ref or call by ref
	r.P = 2 * (float64(r.B) + float64(r.H)) // no need to call (*r).
	return r.P
}

func (r *Square) Area() float64 { // purposely kept r , since r is just an identifer can use any name
	return float64(r.S * r.S)
}

func (s *Square) Perimeter() float64 { // pass by ref or call by ref
	return 4 * float64(s.S)
}

// How to create a constructor ? Go does not follow typical OOPS concepts
