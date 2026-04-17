package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		num1        = 131231
		num2        = uint8(10)
		num3 uint16 = 32334
		num4        = uint32(32343242)
		num5        = 324324.23423
		num6        = float32(2343.34)
	)
	ok1 := true
	str1 := "342423"
	str2 := "3423423.5433"
	any1 := any(float64(num6))
	any2 := any(343.434)
	any3 := any(int(21312))

	var sum float64 // what is the value of sum

	sum = float64(num1) + float64(num2) + float64(num3) + float64(num4) + num5 + float64(num5)

	if ok1 {
		sum += 1
	}

	v1, _ := strconv.Atoi(str1)
	sum += float64(v1)

	v2, _ := strconv.ParseFloat(str2, 64)

	sum += v2

	// v, ok := any1.(float64)

	// if ok {
	// 	sum += v
	// }

	sum = sum + any1.(float64) + any2.(float64) + float64(any3.(int))

	fmt.Printf("%.5f", sum)
}

// go through functions from math package
