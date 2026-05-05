package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {

	r, err := add(true, "hello World")

	fmt.Println(r, err)

	r, err = add(float64(34.34), float64(67.7))

	fmt.Println(r, err)

	r, err = add(float32(123), float32(433.4))
	fmt.Println(r, err)

	r, err = add(int8(123), int8(4))
	fmt.Println(r, err)

	r1 := addG(12, 14)
	println(r1)

	a, b := uint8(10), uint8(16)
	r2 := addG(a, b)
	println(r2)

	r3 := addG(float32(12312.123), float32(123.123))
	println(r3)

}

// type switch
func isNumber(n any) bool {
	switch n.(type) { // n.(type) gives the type of the variable
	case uint, int, uint8, int8, uint16, int16, uint32, int32, uint64, int64, float32, float64:
		return true
	default:
		return false
	}
}

func add(a, b any) (float64, error) {
	sum := 0.0
	if isNumber(a) == true && isNumber(b) == true {
		if reflect.TypeOf(a) != reflect.TypeOf(b) {
			return 0, errors.New("a and b are two different types.I do calc only if both are same types")
		}
		// write the logic here
		switch v := a.(type) {
		case uint8:
			return float64(a.(uint8)) + float64(b.(uint8)), nil
		case uint16:
			return float64(a.(uint16)) + float64(b.(uint16)), nil
		case uint32:
			return float64(a.(uint32)) + float64(b.(uint32)), nil
		case uint64:
			return float64(a.(uint64)) + float64(b.(uint64)), nil
		case uint:
			return float64(a.(uint)) + float64(b.(uint)), nil
			// lets use v for few types
		case int8:
			return float64(v) + float64(b.(int8)), nil
		case int16:
			return float64(v) + float64(b.(int16)), nil
		case int32:
			return float64(v) + float64(b.(int32)), nil
		case int64:
			return float64(a.(int64)) + float64(b.(int64)), nil
		case int:
			return float64(a.(int)) + float64(b.(int)), nil
		case float32:
			return float64(a.(float32)) + float64(b.(float32)), nil
		case float64:
			return a.(float64) + b.(float64), nil
			// default:
			// 	return 0, errors.New("even it is a number, may be not fall under all the number cases.")
		}

	} else {
		return 0, errors.New("invalid type for arth operation")
	}

	return sum, nil
}

func addG[T int | uint | uint8 | int8 | uint16 | int16 | uint32 | int32 | uint64 | int64 | float32 | float64](a, b T) T {
	return a + b
}

func addGI[T INumber](a, b T) T {
	return a + b
}

type INumber interface {
	int | uint | uint8 | int8 | uint16 | int16 | uint32 | int32 | uint64 | int64 | float32 | float64
}
