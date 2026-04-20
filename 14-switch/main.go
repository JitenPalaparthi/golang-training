package main

import "math/rand/v2"

func main() {

	day := rand.IntN(10) // 0-9

	switch day {
	case 1:
		println("Sunday")
	case 2:
		println("Monday")
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Friday")
	case 7:
		println("Saturday")
	default:
		println(day, "---> noday")
	}

	char := rand.IntN(256) // 0-255 are ascii chars

	switch char {
	case 'A', 'E', 'I', 'O', 'U':
		println(string(char), "is an uppercase vovel")
	case 'a', 'e', 'i', 'o', 'u':
		println(string(char), "is an lowecase vovel")
	default:
		println(string(char), "can be cosonent or any char")
	}

	// switch case empty switch

	age, gender := 4, 'm'

	switch {
	case age <= 5 || (gender == 'F' || gender == 'f'):
		println("according to Telangana and AP , it is free to travel in RTC bus")
	default:
		println("need to purchase a ticket to travel ")
	}

	// fallthrough -> where ever you remove break intentionally , you have add fallthrough in go to get the same effect
	num := 12

	switch {
	case num%8 == 0:
		println(num, "is divisible by 8")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	// case num%7 == 0:
	// 	println(num, "is divisible by 7")
	// 	fallthrough
	case num%2 == 0:
		println(num, "is divisible by 2")
	default:
		println(num, "is an odd so not divisible by any one of 8,4,2")
	}

	println("negative fallthrough issues ")

	num = 14
	switch {
	case num%2 == 0:
		println(num, "is divisible by 2")
		fallthrough // fallthrough does not check the case condition, it directly falls to the body of the following case
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%8 == 0:
		println(num, "is divisible by 8")
	default:
		println(num, "is an odd so not divisible by any one of 8,4,2")
	}

}

// no break statement is required, it is auto break
