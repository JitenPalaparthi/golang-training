package main

import (
	r "math/rand"
	"math/rand/v2"
	"shapes/shape"
	"shapes/shape/cuboid"
	"shapes/shape/rect"
	"shapes/shape/square" // module/directory/subdirectory
)

func main() {

	r1 := rect.NewRect(12.43, 14.34)

	r2 := rect.NewRect(87.56, 84.56)

	r3 := rect.NewRect(87.56, 84.56)

	s1 := square.NewSquare(34.5)

	s2 := square.NewSquare(24.5)

	c1 := cuboid.NewCuboid(12.43, 14.56, 6.45)

	shape.Shape(r1)
	shape.Shape(r2)
	shape.Shape(r3)
	shape.Shape(s1)
	shape.Shape(s2)
	shape.Shape(c1)

	println(rand.IntN(100))
	println(r.Intn(1233))

}
