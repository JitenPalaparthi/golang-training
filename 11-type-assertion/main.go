package main

func main() {

	any1 := any(989423432.343) // create any assigned a float64 value

	any1 = uint16(32343) // mutated any1 with uint16 value

	any1 = true

	any1 = "Hello world"

	v1, done := any1.(float64)
	if !done {
		v1, done := any1.(uint16)
		if !done {
			v1, done := any1.(bool)
			if !done {
				println("unable to assert to float64 or uint16 or bool")
			} else {
				println(v1)
			}
		} else {
			println(v1)
		}
	} else {
		println(v1)
	}
}
