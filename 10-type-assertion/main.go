package main

import (
	"math/rand/v2"
	"strconv"
)

func main() {

	num := 100

	var num1 uint8 = uint8(num) // type casting

	str := strconv.Itoa(rand.IntN(9999999)) // fmt.Sprint(rand.IntN(9999999))

	num2, _ := strconv.Atoi(str) // type conversion

	println(num1, num2)

	var any1 any // interface{} // nil , empty interface {}interface

	if any1 == nil {
		println("any1 is nil")
	}

	any1 = num2
	any1 = true

	// var num3 int = any1.(int) // cant do typem casting, type conversion, type assertion
	// // any.(T)

	// num3 := any1.(int)

	num3, done := any1.(int) // the value of the type that has been asserted , bool value gives true or false

	if done == true {
		println(num3)
	} else {
		println("could not be asserted, as expected is int")
	}

	var ok1 bool = any1.(bool)

	println(ok1)

	println("Hello folks .. welcome to panic in Go")
}

// any header
// data ptr
// type ptr
