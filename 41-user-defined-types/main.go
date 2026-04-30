package main

import "fmt"

func main() {

	var p1 struct {
		id    int
		name  string
		email string
	} // not created any type, just the struct as a variable

	p1.id = 101
	p1.name = "Jiten"
	p1.email = "jitenp@outlook.com"

	fmt.Println(p1)

	var p2 struct {
		id    int
		name  string
		email string
	} = struct {
		id    int
		name  string
		email string
	}{100, "jiten", "jitenp@outlook.com"}

	//fmt.Println(p2)

	ShowDetails(p2)

	var p3 struct {
		id     int
		name   string
		email  string
		status string
	} = struct {
		id     int
		name   string
		email  string
		status string
	}{100, "jiten", "jitenp@outlook.com", "active"}

	//ShowDetails(p3) // does not accept. bcz ShowDetails accepts a variable of the type struct {id    int name  string email string}

	fmt.Println(p3)
}

// Where is the structure?

func ShowDetails(p struct {
	id    int
	name  string
	email string
}) {
	fmt.Println("Id:", p.id)
	fmt.Println("Name:", p.name)
	fmt.Println("Email:", p.email)
}
