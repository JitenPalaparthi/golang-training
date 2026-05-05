package main

func main() {

	r1 := NewRect(12.43, 14.34)

	r2 := NewRect(87.56, 84.56)

	s1 := NewSquare(34.5)

	s2 := NewSquare(24.5)

	c1 := NewCuboid(12.43, 14.56, 6.45)

	Shape(r1)
	Shape(r2)
	Shape(s1)
	Shape(s2)
	Shape(c1)

}
