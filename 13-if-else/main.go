package main

import "strconv"

var (
	v3  int
	err error
)

func main() {

	str1 := "4423.2342"

	if v1, err := strconv.ParseFloat(str1, 64); err != nil {
		println(err.Error())
	} else {
		println(v1)
	}

	str2 := "4364545"

	v2, err := strconv.Atoi(str2)

	if err != nil {
		println(err.Error())
	} else {
		println(v2)
	}

	// println(v1)
	println(v2)

	str3 := "7"

	// v3, err = strconv.Atoi(str3)
	// if err != nil {
	// 	println(err.Error())
	// } else {
	// 	println(v3)
	// }

	v, err := GetStrAndSq(str3)

	if err != nil {
		println(err.Error())
	} else {
		println(v)
	}

	println(v3)
}
