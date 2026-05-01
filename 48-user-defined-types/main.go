package main

import "fmt"

func main() {

	var num1 int = 123
	var num2 MyInt1 = 323
	var num3 MyInt2 = 23.2
	var num4 MyInt3 = 4342.34

	//var ok1 bool = true
	//MyInt1(ok1).ToString()

	s := MyInt1(num1).ToString()
	sq := MyInt2(num1).Sq()
	cb := MyInt3(num1).Cube()

	fmt.Println("ToString:", s, "Square:", sq, "Cube:", cb)

	s = num2.ToString() // no need to type cast, num2 is already Myint1
	sq = MyInt2(num2).Sq()
	cb = MyInt3(num2).Cube()
	fmt.Println("ToString:", s, "Square:", sq, "Cube:", cb)

	s = MyInt1(num3).ToString()
	sq = num3.Sq()
	cb = MyInt3(num3).Cube()
	fmt.Println("ToString:", s, "Square:", sq, "Cube:", cb)

	s = MyInt1(num4).ToString()
	sq = MyInt2(num4).Sq()
	cb = num4.Cube()
	fmt.Println("ToString:", s, "Square:", sq, "Cube:", cb)

}

type MyInt1 int

type MyInt2 float32

type MyInt3 float64

func (m MyInt1) ToString() string {
	return fmt.Sprint(m)
}

func (m MyInt2) Sq() float64 {
	return float64(m * m)
}

func (m MyInt3) Cube() float64 {
	return float64(m * m * m)
}
