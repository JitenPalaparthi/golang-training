package main

const (
	PI  float32 = 3.14
	MAX         = 9999
)

var Count = 999999 // Global variable, package level variable , static variables

var Global int // unassigned, zero value

func main() {
	str := "Hello, Golang Minds!"
	println(str)

	var a, b = 10, 20 // stack memory

	println(a, b)

	//PI = 3.145 // it does not accept why it is a constant
}

// strings
// any
