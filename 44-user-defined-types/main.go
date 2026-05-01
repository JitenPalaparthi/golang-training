package main

import "fmt"

func main() {

	r1 := Rect{10.2, 13.5, 0, 0}

	a1 := r1.Area()
	p1 := r1.Perimeter()

	fmt.Printf("Area of Rect:%.2f\nPerimeter of Rect:%.2f\n", a1, p1)

	fmt.Printf("Area of Rect:%.2f\nPerimeter of Rect:%.2f\n", r1.A, r1.P)

	a2 := r1.AreaP()         // r1 is not a pointer still can call a method through a pointer receiver, go does the work
	p2 := (&r1).PerimeterP() // mo need to do this .. go does

	fmt.Printf("Area of Rect:%.2f\nPerimeter of Rect:%.2f\n", a2, p2)
	fmt.Printf("Area of Rect:%.2f\nPerimeter of Rect:%.2f\n", r1.A, r1.P)

	r2 := &Rect{B: 10.4, H: 12.5}

	a3 := r2.AreaP()
	p3 := r2.PerimeterP()

	fmt.Printf("Area of Rect:%.2f\nPerimeter of Rect:%.2f\n", a3, p3)
	fmt.Printf("Area of Rect:%.2f\nPerimeter of Rect:%.2f\n", r2.A, r2.P)

}

type Rect struct {
	B, H float32
	A, P float64
}

// Area and Perimeter , two are functions

func Area(r Rect) float64 { // pass by value or call by value
	r.A = float64(r.B) * float64(r.H)
	return r.A
}

func Perimeter(r Rect) float64 { // pass by value or call by value
	r.P = 2 * (float64(r.B) + float64(r.H))
	return r.P
}

// The below are methods

func (r Rect) Area() float64 { // pass by value or call by value
	r.A = float64(r.B) * float64(r.H)
	return r.A
}

func (r Rect) Perimeter() float64 { // pass by value or call by value
	r.P = 2 * (float64(r.B) + float64(r.H))
	return r.P
}

func (r *Rect) AreaP() float64 { // pass by ref or call by ref
	r.A = float64(r.B) * float64(r.H)
	return r.A
}

func (rect *Rect) PerimeterP() float64 { // pass by ref or call by ref
	(*rect).P = 2 * (float64(rect.B) + float64(rect.H)) // no need to call (*r).
	return rect.P
}

// Two types of receivers

// Pointer receiver and non pointer receiver

// Pointer receiver is call by ref and non pointer receiver is call by value

// 99% receivers are pointer receivers

// just becase the receiver is a pointer receiver the variable/object not needs to be a pointer

// how many receivers can a method have ? only one

// general convenction is create a receiver with single char (r *Rect) does not mean cannot give multiple chars

// For the same type receiver names can be different for different methods , no issues
