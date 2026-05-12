package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, 世界", time.Now())

	var (
		num1, num2 int = 100, 200
		num3       int = num1 + num2
	)

	// arth: +,-,*,/,%
	// var num1 int = 100
	// var num2 int = 200
	// var num3 int = num1 + num2 // basic arth operation
	// Expression

	fmt.Println("Sum of num1 and num2:", num3)

	// Zero values

	var (
		num4   int     // zero value is for numbers 0
		ok1    bool    // zero value for bool is false
		str1   string  // zero value for string is ""
		float1 float32 // zero value is for numbers 0
	)

	fmt.Println("num4:", num4, "ok1:", ok1, "str1:", str1, "float1:", float1)

	num4 = num3 * 10 // mutation

	fmt.Println("num4:", num4)

	// type inference
	var (
		num5   = 12             // infer it as int
		float2 = 989.34         // infer it as float64
		ok2    = true           // infer it as bool
		str2   = "Hello World!" // infer it as string
	)

	fmt.Println("num5:", num5, "float2:", float2, "ok2:", ok2, "str2:", str2)

	// int8,uint8

	/*
		var (
			num6 uint8  = 100
			num7 uint8  = 200
			num8 uint16 = uint16(num6) + uint16(num7) // 32xxx
		)
		fmt.Println("num8:", num8)
	*/

}

// value inference
//
