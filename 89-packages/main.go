package main

import (
	"shapes/cuboid"
	r "shapes/rectangle"         // alias
	rect "shapes/rectangle"      // alias
	rectangle "shapes/rectangle" // alias
	"shapes/shape"
	"shapes/square"
)

func main() {

	r1 := rectangle.NewRect(12.43, 14.34)

	r2 := rect.NewRect(87.56, 84.56)

	r3 := r.NewRect(87.56, 84.56)

	s1 := square.NewSquare(34.5)

	s2 := square.NewSquare(24.5)

	c1 := cuboid.NewCuboid(12.43, 14.56, 6.45)

	shape.Shape(r1)
	shape.Shape(r2)
	shape.Shape(r3)
	shape.Shape(s1)
	shape.Shape(s2)
	shape.Shape(c1)

}
