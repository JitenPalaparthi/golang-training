package main

func main() {

	//c1 := ColourCode{9999, "Red", "Full Redf"}

	c1 := ColourCode{9999, "Red"}

	println("Code:", c1.int, "Colour:", c1.string)

}

// Type with anonymous fileds
type ColourCode struct {
	int
	string
	//SColour
}

//type SColour = string // alias
