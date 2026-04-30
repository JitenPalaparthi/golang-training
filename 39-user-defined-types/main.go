package main

import "fmt"

func main() {

	p1 := Person{Id: 101, Name: "Jiten", Email: "JitenP@Outlook.Com", Mobile: "9618558500", Address: Address{City: "Guntur", PinCode: "522001", Status: "Current"}}

	fmt.Println("Name:", p1.Name, "City:", p1.Address.City)

	e1 := Employee{Id: 101, Name: "Jiten", Email: "JitenP@Outlook.Com", Mobile: "9618558500", Status: "active", Address: Address{City: "Guntur", PinCode: "522001", Status: "Current"}}

	fmt.Println("Name:", e1.Name, "Status:", e1.Status, "City:", e1.City, "Addres Status:", e1.Address.Status)

}

type Person struct {
	Id      int
	Name    string
	Email   string
	Mobile  string
	Address Address // Composition
}

type Address struct {
	City    string
	PinCode string
	Status  string
}

// Go there is no inheritance

// promoted fields

type Employee struct {
	Id      int
	Name    string
	Email   string
	Mobile  string
	Status  string
	Address // promoted field
}
