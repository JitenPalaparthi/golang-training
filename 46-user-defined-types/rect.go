package main

type Rect struct {
	B, H float32
	A, P float64
}

// new is a function can think it as a constructor
func NewRect(b, h float32) *Rect {
	return &Rect{B: b, H: h}
}

func (r *Rect) Area() float64 { // pass by ref or call by ref
	r.A = float64(r.B) * float64(r.H)
	return r.A
}

func (r *Rect) Perimeter() float64 { // pass by ref or call by ref
	r.P = 2 * (float64(r.B) + float64(r.H)) // no need to call (*r).
	return r.P
}
