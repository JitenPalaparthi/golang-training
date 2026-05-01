package main

import "fmt"

func main() {

	r1 := Rect{10.2, 13.5, 0, 0}

	a1 := Area(r1)
	p1 := Perimeter(r1)

	fmt.Printf("Area of Rect:%.2f\nPerimeter of Rect:%.2f\n", a1, p1)

	fmt.Printf("Area of Rect:%.2f\nPerimeter of Rect:%.2f\n", r1.A, r1.P)

	a2 := AreaP(&r1)
	p2 := PerimeterP(&r1)

	fmt.Printf("Area of Rect:%.2f\nPerimeter of Rect:%.2f\n", a2, p2)
	fmt.Printf("Area of Rect:%.2f\nPerimeter of Rect:%.2f\n", r1.A, r1.P)

	r2 := &Rect{B: 10.4, H: 12.5}

	a3 := AreaP(r2)
	p3 := PerimeterP(r2)

	fmt.Printf("Area of Rect:%.2f\nPerimeter of Rect:%.2f\n", a3, p3)
	fmt.Printf("Area of Rect:%.2f\nPerimeter of Rect:%.2f\n", r2.A, r2.P)

}

type Rect struct {
	B, H float32
	A, P float64
}

func Area(r Rect) float64 { // pass by value or call by value
	r.A = float64(r.B) * float64(r.H)
	return r.A
}

func Perimeter(r Rect) float64 { // pass by value or call by value
	r.P = 2 * (float64(r.B) + float64(r.H))
	return r.P
}

func AreaP(r *Rect) float64 { // pass by ref or call by ref
	r.A = float64(r.B) * float64(r.H)
	return r.A
}

func PerimeterP(r *Rect) float64 { // pass by ref or call by ref
	(*r).P = 2 * (float64(r.B) + float64(r.H)) // no need to call (*r).
	return r.P
}
