package main

import (
	"errors"
	"fmt"
)

func main() {

	m1 := make(MyMap)

	m1["name"] = "Jiten"
	m1["age"] = 43
	m1["ismarried"] = true
	m1["gender"] = 'M'
	m1["address"] = "Trivandrum, 695401"

	if err := m1.Delete("email"); err != nil {
		println(err.Error())
	} else {
		println("email successfully deleted from the map")
	}

	if err := m1.Delete("age"); err != nil {
		println(err.Error())
	} else {
		println("age successfully deleted from the map")
	}

	//fmt.Println(m1)
	m1.Show()

	m2 := make(map[string]any) // normal map not user defined map

	m2["name"] = "Jiten"
	m2["age"] = 43
	m2["ismarried"] = true
	m2["gender"] = 'M'
	m2["address"] = "Trivandrum, 695401"

	if err := MyMap(m2).Delete("email"); err != nil {
		println(err.Error())
	} else {
		println("email successfully deleted from the map")
	}

	if err := MyMap(m2).Delete("age"); err != nil {
		println(err.Error())
	} else {
		println("age successfully deleted from the map")
	}

	MyMap(m2).Show()
	clear(m2) // can call clear func even clear is only for slice and map but m2 is MyMap which is based on map[string]any can call all builtin functions

	// type cast works only if the underlining maps are same

	// uncomment the below code to run and check

	// m3 := make(map[string]string)
	// m3["name"] = "Jiten"
	// m3["age"] = "43"
	// m3["ismarried"] = "true"
	// m3["gender"] = "M"
	// m3["address"] = "Trivandrum, 695401"

	// MyMap(m3).Show() // it does not work bcz MyMap is map[string]any but m3 is map[string]string

}

type MyMap map[string]any

func (m MyMap) Delete(key string) error {
	if m == nil {
		return errors.New("nil map")
	}
	_, ok := m[key]
	if !ok {
		return fmt.Errorf("key:%v does not exist", key)
	}
	delete(m, key) // it deletes
	return nil
}

func (m MyMap) Show() {
	for k, v := range m {
		fmt.Println("Key:", k, "Value:", v)
	}
}

// even for the user defined types on maps, slices or even any other types
// as long as the built in functions work on the direct type , they can also work on the user defined types
// range , make, append , clear, copy
