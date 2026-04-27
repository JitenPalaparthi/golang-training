package main

import (
	"errors"
	"fmt"
)

func main() {

	var map1 map[string]string // only declared but not instantiated

	if map1 == nil {
		println("nil map")
		map1 = make(map[string]string)
	}
	// map1["name"] = "Jiten"
	// map1["age"] = "42"
	// map1["address"] = "Trivandrum Medical College"

	map1["560086"] = "Bangalore-1"
	map1["560096"] = "Bangalore-2"
	map1["645901"] = "Trivandrum-1"
	map1["522001"] = "Guntur-1"
	map1["00000"] = "" // This is a default value w.r.t string

	map2 := make(map[string]string) // no elements in the map but shorthand declaration

	map2["560086"] = "Bangalore-1"
	map2["560096"] = "Bangalore-2"
	map2["645901"] = "Trivandrum-1"
	map2["522001"] = "Guntur-1"

	map3 := map[string]string{"560086": "Bangalore-1", "560096": "Bangalore-2"} // not using make but creating and instantiate a map directly with values

	for k, v := range map1 {
		println("Key:", k, "Value:", v)
	}

	for k, v := range map2 {
		println("Key:", k, "Value:", v)
	}

	for k, v := range map3 {
		println("Key:", k, "Value:", v)
	}

	// can also get a value from the map

	v := map1["560086"]

	println("Fetching value by using key:", v)

	v, ok := map1["560091"] // value,ok variation, ok is bool, if true there is a value for that key else key does not exist and hence no value
	if !ok {
		println("key does not exist")
	} else {
		println("Fetching value by using key:", v)
	}
	println()

	if v, err := GetValue(map1, "560091"); err != nil {
		println(err.Error())
	} else {
		println("Fetching value by using GetValue func:", v)

	}

	println("after delete")

	delete(map1, "560086") // it deletes only if map is not nill and key exists
	// either map is nil or key does not exist delete does nothing

	v, ok = map1["560086"] // value,ok variation, ok is bool, if true there is a value for that key else key does not exist and hence no value
	if !ok {
		println("key does not exist")
	} else {
		println("Fetching value by using key:", v)
	}

	println("delete does nothing below")

	var map4 map[string]string // nil map
	delete(map4, "12312")      // it does nothing why map4 is nil, still delete does not give any kind fo information

	println("clear clears the map")

	clear(map3) // all elements in the map are deleted

	fmt.Println(map3)
	//fmt.Println(map2)

}

// No need to do this, to iterate all elemernts to find a key
func GetValue(m map[string]string, key string) (string, error) {
	if m == nil {
		return "", errors.New("nil map")
	}

	ok := false
	for k, v := range m {
		if k == key {
			ok = true
			return v, nil
		}
	}

	if !ok {
		return "", errors.New("key does not exist")
	}
	return "", nil

}

// What is map -> pair of keys and values
// map can be nil
// to instantiate a map can use make builtin function
// can also give size while making a map using make function

// to iterate through the map use range loop
// map is not ordered, and also not thread safe
// what can be a key in a map --> any datatype that does == operation can be a key

// sha --> Secured Hashing Algo
// 648a6a6ffffdaa0badb23b8baf90b6168dd16b3a
// 648a6a6ffffdaa0badb23b8baf90b6168dd16b3a
