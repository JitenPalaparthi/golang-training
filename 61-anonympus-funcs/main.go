package main

import "errors"

func main() {

	func() {
		println("Hello World")
	}()

	slice := []int{10, 23, 4, 1, 25, 6, 5, 33, 57, 4, 54, 8, 564, 34, 667, 565}

	max, err := func(slice []int) (int, error) {
		if slice == nil {
			return 0, errors.New("nil slice")
		}
		if len(slice) > 0 {
			max := slice[0]

			for _, v := range slice {
				if v > max {
					max = v
				}
			}
			return max, nil
		}
		return 0, errors.New("empty slice")
	}(slice)

	if err != nil {
		println(err.Error())
	} else {
		println("max:", max)
	}

	// min
	minfn := func(slice []int) (int, error) {
		if slice == nil {
			return 0, errors.New("nil slice")
		}
		if len(slice) > 0 {
			min := slice[0]

			for _, v := range slice {
				if v < min {
					min = v
				}
			}
			return min, nil
		}
		return 0, errors.New("empty slice")
	} // no executor

	// minfn is a function

	min, err := minfn(slice)

	if err != nil {
		println(err.Error())
	} else {
		println("min:", min)
	}

}
