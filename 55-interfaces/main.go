package main

func main() {

	r1 := NewRect(12.43, 14.34)

	r2 := NewRect(87.56, 84.56)

	s1 := NewSquare(34.5)

	s2 := NewSquare(24.5)

	c1 := NewCuboid(12.43, 14.56, 6.45)

	ishapes := make([]IShape, 0)

	//ishapes := make([]any, 0)

	ishapes = append(ishapes, r1, r2, s1, s2, c1, NewCuboid(12.12, 13.14, 15.16), NewSquare(123.43), NewRect(12.34, 98.67))

	//ishapes = append(ishapes, 12313) // this can be doen since ishapes is any
	Shapes(ishapes)

}

func Shapes(ishapes []IShape) {
	for _, shape := range ishapes {
		Shape(shape)
	}
}

// func Shapes(ishapes []any) {
// 	for _, shape := range ishapes {
// 		switch ishape := shape.(type) {
// 		case IShape:
// 			Shape(ishape)
// 		default:
// 			println("It is not a valid type so not performing any action")
// 		}
// 	}
// }
