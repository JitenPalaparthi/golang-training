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

	//c := new(shape.circle) // cannot be created,it is restricted type

	cr1 := new(shape.Circle) // cannot be created,it is restricted type
	cr1.R = 123.32
	println("Circle Radious", cr1.R)
	//println("Circle a", cr1.a) // a is lowrcase
	println("Circle P", cr1.P) // a is lowrcase
	//shape.greet()            // g starts with lowecase so it is restricted
	shape.Greet()

	cr2 := shape.Newcirlce(232.34)
	println("R", cr2.R)
	//println("a and p",cr2.a,cr2.p) // cant access
}

// Go does not have public, private, protected, internal etc..
// Go package management and access is pretty simple
// With in a package, the type or function or even a valiable or anything in package ,
// if they are created with Uppercase, they are called as Unrestricted, Exported --> Public
// If they are created with lowerCase they are called as Restricted or unexported -> Private
