package main

func main() {
	// 	count := 1
	// loop:
	// 	println(count)
	// 	count++
	// 	if count <= 10 {
	// 		goto loop
	// 	}

	num := 0
loop:
	num++
	if num > 10 {
		goto out
	}
	if num%2 == 0 {
		goto even
	} else {
		goto odd
	}

even:
	println(num, "is even")
	goto loop
odd:
	println(num, "is odd")
	goto loop
out:
	println("exiting the goto blocks")

	println("It is a normal execution")
}
