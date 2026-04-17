package main

import "strconv"

func GetStrAndSq(str string) (int, error) {
	v3, err = strconv.Atoi(str)
	if err != nil {
		return 0, err
	} else {
		v3 = v3 * v3
		return v3, nil
	}
}
