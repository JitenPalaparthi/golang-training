package main

func main() {

	str := "Hello World" // what is string , string is collection of ?

	for i, v := range str {
		println(i, string(v))
	}

	str = "Hello, 世界"

	for i := 0; i < len(str); i++ {
		print(string(str[i]), "")
	}
	println()

	for i, v := range str {
		println(i, string(v))
	}

	println()

	for _, v := range str {
		println(string(v))
	}

	// 1.22 onwards

	for i := range 10 {
		print(i, " ")
	}
	println()

}

// range can be used on arrays, strings, slices, maps and channels
// on arrays, strings, slices,range gives index and value
// on maps range gives key and value
// on channel range gives a value received from a channel
