// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// var a = 100 // type inference
	// var b = 200 // type inference

	/*
		var (
			a, b = 100, 200
		)
	*/

	// shorthand notation, no need to use var, and the type
	// a:= 100 // var a int = 100
	// b:= 200 // var b int = 200

	a, b := 100, 200 // group of shorthand variables

	println("a:", a, "b:", b)

	a1, b1, f1, ok1, str1 := 100, 200, 3.14, true, "Hello World!"
	fmt.Println("a1:", a1, "b1:", b1, "f1:", f1, "ok1:", ok1, "str1:", str1)
	// reflect is package --> reflect.TypeOf

	fmt.Println("a1 type:", reflect.TypeOf(a1), "b1 type:", reflect.TypeOf(b1))
	fmt.Println("ok1 type", reflect.TypeOf(ok1), "f1 type:", reflect.TypeOf(f1), "str1 type:", reflect.TypeOf(str1))
	c1 := (a1+b1)*3 + 10/2 + 100 + (b1 - a1) // arth
	// 300 * 3 + 10/2+100+(b1-a1)
	// 300 * 3 + 10/2+100+100
	// 900 +10/2+100+100
	// 900+5+100+100
	// 905+100+100
	// 1005+100
	// 1105
	fmt.Println("c1:", c1)
}

// (), * /, +,-
