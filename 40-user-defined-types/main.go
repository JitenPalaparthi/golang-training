package main

import "fmt"

func main() {

	p1 := Person{Id: 101, Name: "Jiten", Email: "JitenP@Outlook.Com", Mobile: "9618558500", Status: "active", Address: struct {
		City    string
		PinCode string
		Status  string
	}{City: "Guntur", PinCode: "522001", Status: "current"}}

	fmt.Println("Name:", p1.Name, "City:", p1.Address.City)
	fmt.Println(p1)

}

type Person struct {
	Id      int
	Name    string
	Email   string
	Mobile  string
	Status  string
	Address struct { // Embedded struct
		City    string
		PinCode string
		Status  string
	}
}
