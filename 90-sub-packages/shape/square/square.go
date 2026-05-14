package square

type Square float64

func NewSquare(s float64) Square {
	return Square(s)
}

func (s Square) Area() float64 {
	return float64(s * s)
}

func (s Square) Perimeter() float64 {
	return float64(4 * s)
}
