package main

func main() {

	/// What datatype can be  a key

	arr1 := [2]int{10, 20}
	arr2 := [2]int{20, 10}

	if arr1 == arr2 {
		println("both are equal")
	} else {
		println("both are not equal")
	}

	// since can perform == operation on arrays

	map1 := make(map[[2]int]string) // [2]int can be a key, not only [2]int any array type can be a key

	map1[arr1] = "some values in the arr1"
	map1[arr2] = "some values in the arr2"

	// can a slice be a key?

	// slice1 := []int{10, 20}
	// slice2 := []int{20, 10}

	// if slice1 == slice2 { // == is an invalid operation on slice, so a slice cannot be used as a key type to the map
	// 	println("both are equal")
	// } else {
	// 	println("both are not equal")
	// }

	// since can perform == operation on arrays

	// map2 := make(map[[]int]string)

	// map2 := make(map[string][]any) --> value can be any valid type no issues, but key should be only a datatype that accepts == operation

	// bool can be a key though can give only two values

	map2 := make(map[bool]string)

	map2[true] = "Something is true"
	map2[false] = "Something is false"

	any1 := any(100)
	any2 := any("Hello World")

	if any1 == any2 { // it accepts to perform == operation so any can be a key

	}

	map3 := make(map[any]any)

	map3[any1] = "Some ing"
	map3[any2] = "Some string"
}
