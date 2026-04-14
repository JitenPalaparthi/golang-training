package main

import (
	"fmt"
	"math/rand/v2"
	"reflect"
)

func main() {

	// var num int = 123312

	// fmt.Println("Type of num:", reflect.TypeOf(num))

	var a1 any // interface{}

	fmt.Println("data:", a1, "type:", reflect.TypeOf(a1)) // ?

	// var str2 string
	// fmt.Println(str2, "-->")

	// What is the type here?

	a1 = 100                                              // int
	fmt.Println("data:", a1, "type:", reflect.TypeOf(a1)) // ?

	a1 = true                                             // bool
	fmt.Println("data:", a1, "type:", reflect.TypeOf(a1)) // ?

	a1 = 534534.345                                       // float64
	fmt.Println("data:", a1, "type:", reflect.TypeOf(a1)) // ?

	str1 := "hello World!"

	a1 = str1                                             // string
	fmt.Println("data:", a1, "type:", reflect.TypeOf(a1)) // ?

	var a2 any = rand.IntN(3434324)

	a1 = a2                                               // any
	fmt.Println("data:", a1, "type:", reflect.TypeOf(a1)) // ?

	var num1 uint8 = 145

	a1 = num1
	fmt.Println("data:", a1, "type:", reflect.TypeOf(a1)) // ?

	// a1 = Person{ID: 12312}

}

// Go is type safe | statically typed programming langauge

// any also has a header

// any header
// data pointer
// type pointer , all the types are loaded in memory, that address it assignes

// When ever you see a pointer , think of nil except string header

// There is no null in Go , only nil
// nil is used where ever there are pointers , directly or indirecly
