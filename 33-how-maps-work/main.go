package main

import (
	"fmt"
	"hash/maphash"
)

var seed = maphash.MakeSeed()

func main() {

	hash := Hash64("Jiten")
	fmt.Printf("0x%x\n", hash)
	fmt.Printf("Full Hash: %b\n", hash)
	// 1110010011011101000011001101100001111110001000101100100110000111
	// H1 and H2 which are higher bit and lowerbit

	//map1 := map[string]string{"Jiten": "961812312312"}

	h1 := hash >> 7 // 0000000111001001101110100001100110110000111111000100010110010011
	fmt.Printf("top hash:  %b\n", h1)

	//h2 := byte(hash & 0x7F) // masking
	h2 := byte(hash & 0b1111111) // masking
	fmt.Printf("lower hash:%b\n", h2)

	// 111001001101110100001100110110000111111000100010110010011 0000111
	// 000000000000000000000000000000000000000000000000000000000 1111111
	// 000000000000000000000000000000000000000000000000000000000 0000111

	// shifting to the right
	var num uint = 21321312

	fmt.Printf("%b\n", num) // 1010001010101011001100000

	num1 := num >> 4 // 0000101000101010101100110

	fmt.Printf("%b\n%d\n", num1, num1)

	// num2 := num << 4

	// fmt.Printf("%b\n%d\n", num2, num2)

}

func Hash64(s string) uint64 {
	return maphash.String(seed, s)
}
