package main

import "math/rand/v2"

func main() {

	println("notmal loop")
	for i := 1; i <= 10; i++ {
		print(i, " ")
	}
	println()

	println("for looks like while")
	a, b, count := 0, 1, 1

	for count <= 10 { // for loop acts like while loop
		print(a, " ")
		a, b = b, a+b
		count++
	}
	println()

	num := 1
	println("for without init")
	for ; num <= 10; num++ {
		if num%2 == 0 {
			continue // continue skips the current iteration
		}
		print(num, " ")
	}

	println()

	num = 1
	println("for unconditional, using break")
	for {
		if num > 10 {
			break // breaks the loop , switch and for
		}
		if num%2 == 0 {
			print(num, " ")
		}
		num++
	}

	println()
	println("for multipule init, and bigger condition as an expression")
	for i, j := 1, 10; (i <= 5 && j >= 5) && (true || false); i, j = i+1, j-1 {
		println("i:", i, "j:", j, "condition:", (i <= 5 && j >= 5) && (true || false))
	}

	println("nested loop")

	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			println("i:", i, "j:", j)
		}
	}

	println("nested loop with break")

	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if j == 3 {
				break // it breaks j , why break is written in j loop but not i loop
			}
			println("i:", i, "j:", j)
		}
	}

	// println("break all loops")
	// done1, done2 := false, false
	// for i := 1; i <= 5; i++ {
	// 	if done1 {
	// 		break
	// 	}
	// 	for j := 1; j <= 5; j++ {
	// 		if done2 {
	// 			break
	// 		}
	// 		for k := 1; k <= 5; k++ {
	// 			if k == 3 {
	// 				done1, done2 = true, true
	// 				break
	// 			}
	// 			println("i:", i, "j:", j, "k:", k)
	// 		}
	// 	}
	// }

	println("break all loops")

outer:
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			for k := 1; k <= 5; k++ {
				if k == 3 {
					break outer // break with a label
				}
				println("i:", i, "j:", j, "k:", k)
			}
		}
	}

	println("using break label in switch")

	// until the number is divisible by 17
out:
	for {
		num = rand.IntN(9999)
		switch {
		case num%8 == 0:
			println("even number, divisible by 8:", num)
			fallthrough
		case num%4 == 0:
			println("even number, divisible by 4:", num)
			fallthrough
		case num%2 == 0:
			println("even number, divisible by 2:", num)
			fallthrough
		default:
			if num%17 == 0 {
				println(num, "is divisible by 17 so exiting")
				break out
			}
		}
	}

}
