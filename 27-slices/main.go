package main

import "fmt"

func main() {

	var slice1 []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice1)
	slice2 := slice1 // what gets copied, the headers are copied

	slice2[0] = 11111
	fmt.Println(slice1) // change of slice2[0] would also change slice1

	slice2 = append(slice2, 11) // append slice2

	slice2[1] = 2222
	fmt.Println(slice1)

	slice3 := make([]int, 10, 20)
	// slice3 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice3[0], slice3[1], slice3[2], slice3[3], slice3[4], slice3[5], slice3[6], slice3[7], slice3[8], slice3[9] = 1, 2, 3, 4, 5, 6, 7, 8, 9, 10

	//fmt.Println(slice3)

	slice4 := slice3
	slice4[0] = 1111
	fmt.Println("slice3", slice3)
	slice4 = append(slice4, 11)

	slice4[1] = 2222
	fmt.Println("slice3", slice3)

	// copy

	slice5 := make([]int, 5)

	copy(slice5, slice1) // deep copy, the values are copied both are different slices different pointers etc

	slice5[0] = 99999

	// since it is a deep copy both of them are two different slcies, no change happen when one of the slices are changed
	fmt.Println("slice1", slice1)
	fmt.Println("slice5", slice5)

	// clear

	clear(slice5) // clear does not make a slice nil rather it makes all the values are zero values
	fmt.Println("slice5", slice5)

	strslcie := []string{"Jiten", "Farzeen", "Vishnu", "Sree", "Sai"}

	fmt.Println(strslcie)
	clear(strslcie) // what is the zero value of a string ""
	fmt.Println(strslcie)

}

// we dont deallocate a slice, it has to happen automatically by the GC or based on the scope
