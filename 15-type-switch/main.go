package main

func main() {

	ok := isNumber(true)

	println(ok)

	ok = isNumber("Hello World")
	println(ok)

	ok = isNumber(float32(213.23))
	println(ok)

	ok = isNumber(int64(999999999999999999))
	println(ok)

	ok = isNumber(uint64(9999999999999999999))
	println(ok)

	ok = isNumber2(uint64(9999999999999999999))
	println(ok)

}

// type switch
func isNumber(n any) bool {
	switch n.(type) { // n.(type) gives the type of the variable
	case uint, int, uint8, int8, uint16, int16, uint32, int32, uint64, int64, float32, float64:
		return true
	default:
		return false
	}
}

// can use val,bool kind of assertion and also check but it is quite difficult or confusing
func isNumber2(n any) bool {
	if _, ok := n.(int); !ok {
		if _, ok := n.(uint); !ok {
			if _, ok := n.(uint8); !ok {
				if _, ok := n.(int8); !ok {
					if _, ok := n.(uint16); !ok {
						if _, ok := n.(int16); !ok {
							if _, ok := n.(uint32); !ok {
								if _, ok := n.(int32); !ok {
									if _, ok := n.(uint64); !ok {
										if _, ok := n.(int64); !ok {
											if _, ok := n.(float32); !ok {
												if _, ok := n.(float64); !ok {
													return false
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		return true
	}

	return true
}
