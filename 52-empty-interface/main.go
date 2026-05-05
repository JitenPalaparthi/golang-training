package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	var e1 Empty = int64(2213121)
	var a1 any = "Hello World"
	var i1 interface{} = true

	fmt.Println(e1, a1, i1)
	var num1 = 43434
	sum := num1 + int(e1.(int64))
	println("sum = ", sum)

	e1 = a1 // data pointer and type pointer are assigned to the e1 struct internally

	fmt.Println("Type of e1:", reflect.TypeOf(e1), "size of e1:", unsafe.Sizeof(e1))

	var e2 Empty // like any it is also nil if no value is assigned to it

	if e2 == nil {
		fmt.Println("Type of e2:", reflect.TypeOf(e2), "Value of e2:", e2, "size of e2:", unsafe.Sizeof(e2))
	}

}

type Empty interface{} // type ptr and aslo data ptr
