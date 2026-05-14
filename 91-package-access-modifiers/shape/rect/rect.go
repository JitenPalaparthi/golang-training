package rect

import "shapes/shape"

type Rect struct {
	L, B float32
}

func NewRect(l, b float32) *Rect {
	return &Rect{l, b}
}

func (r *Rect) Area() float64 {
	return float64(r.L) * float64(r.B)
}

func (r *Rect) Perimeter() float64 {
	return 2 * float64(r.L+r.B)
}

func Greet() {
	println("Welcome to Greet function from rect package")
	shape.Greet() // To call a functuon from the main package into the sub package.. no special keywords like super etc..
	//shape.greet() // cannot access still greet is only be accessed in shape pacakge, even rect is a sub package the rules are same, no protected kind of
}
