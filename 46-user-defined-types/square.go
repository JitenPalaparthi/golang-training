package main

func NewSquare(s float32) *Square {
	return &Square{s}
}

func (r *Square) Area() float64 { // purposely kept r , since r is just an identifer can use any name
	return float64(r.S * r.S)
}

func (s *Square) Perimeter() float64 { // pass by ref or call by ref
	return 4 * float64(s.S)
}

type Square struct {
	S float32
}

// order does not matter
