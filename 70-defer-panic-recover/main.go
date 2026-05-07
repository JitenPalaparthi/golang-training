package main

import (
	"fmt"
	"os"
)

func main() {

	defer println("Hello World, end of Main")
	defer recoverThis()

	n, err := Write("data.txt", "Hello World, How are you doing")
	if err != nil {
		println(err.Error())
	} else {
		if n != 0 {
			println("data successfully stored in the file")
		}
	}

	defer func() {
		a, b := 0, 1
		for range 10 {
			print(a, "  ")
			a, b = b, a+b
		}
	}()
	defer println("Hello World")

	num, div := 100, 0
	println(num / div) // it would be a panic
}

func Write(filename string, data string) (int, error) {
	// var err error
	defer func() {
		println("Write function is called , whether panic or not I am getting executed")
	}()
	defer recoverThis()
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err.Error()) // user defined panic, custom panic , when ever there is a panic before panic,
		//return 0, err
	}

	defer func() {
		println("There is a panic above, would I be executed?") // No
	}()
	defer f.Close()
	//f.Close()
	return f.Write([]byte(data))
	//return n, err
}

// recover, recovers from a panic, do not crash the whole call stack, rather limit the panic only to that function or caller

func recoverThis() {
	if r := recover(); r != nil { // recover function is a built in function, what it gives nil or panic data
		fmt.Println(r)
		// write the logic to do what ever during panic
	}
	// } else {
	// 	// do nothing
	// }
}
