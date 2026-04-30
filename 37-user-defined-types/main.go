package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//var num1 int

	var p1 Person // p1 is a variable of a Person type

	p1.Id = 100
	p1.Name = "Jiten"
	p1.Email = "JitenP@Outlook.Com"
	p1.Mobile = "9618558500"

	fmt.Println(p1)
	fmt.Println("Id:", p1.Id, "Name:", p1.Name, "Email:", p1.Email, "Mobile:", p1.Mobile)

	// assigning values without filed names, so all values are automatically set to the respective fileds, but order of assignment is very important
	p2 := Person{101, "Jiten", "JitenP@Outlook.Com", "9618558500"}
	fmt.Println(p2)

	p3 := Person{} // Dont give any values ..zero values are given

	fmt.Println(p3)

	p4 := Person{Mobile: "9618558500", Email: "JitenP@Outlook.Com", Name: "Jiten", Id: 100}
	fmt.Println(p4)

	p5 := &Person{101, "Jiten", "JitenP@Outlook.Com", "9618558500"} // The memory is allocated and that address is given to p5

	fmt.Println("Size of p5:", unsafe.Sizeof(p5))

	(*p5).Email = "Jiten.Palaparthi@Gmail.Com" // Dereference the pointer and access the field

	p5.Email = "JitenP@Gmail.Com" // can do like this

	p6 := new(Person) // new allocates memory and assigns that address to the variable so p6 is a pointer of Person

	p6.Id = 111
	p6.Name = "JP"
	p6.Email = "jp@spanlet.com"
	p6.Mobile = "9618558500"

	fmt.Println(p6)
}

// There is not class, there is struct

// This is a struct and is a user defined type
type Person struct {
	Id     int
	Name   string
	Email  string
	Mobile string
}
