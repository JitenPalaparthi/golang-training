package main

func main() {

	num, ptrnum1, ptrnum2 := 100, new(100), new(100)

	defer println("defer num", num, "ptrnum1:", *ptrnum1) // the defer func maintains its stack, so the args, call by value
	defer incr(ptrnum2)
	defer println() // call by reference

	num = num + 1 // num++ or num+=1
	*ptrnum1 += 1
	incr(ptrnum2)

	println("normal num", num, "ptrnum:", *ptrnum1)
	str := "Hello World"
	for _, v := range str {
		defer print(string(v))
	}

}

func incr(ptr *int) {
	if ptr != nil {
		*ptr += 1
		println("ptr:", *ptr)
	}
}
