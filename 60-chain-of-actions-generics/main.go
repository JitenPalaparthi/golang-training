package main

func main() {

	// rows := db.query("users").where("status=active").limit(10).offset(10).exec() //chain of actions, it is just example but not really similar query there in Go

	c1 := New[int](10).Add(10).Mul(2).Add(10).Sub(10).Div(2).Mul(4).Div(1).Add(2).Mul(4).Get()
	println(c1)

	c2 := New[float64](10.123).Add(10).Mul(2.43).Add(10.243).Sub(10.3).Div(2).Mul(4.5).Div(1.3).Add(2).Mul(4).Get()
	println(c2)

}

type Calc[T INumber] struct { // a stru t with generic parameter T
	data T
}

func New[T INumber](d T) *Calc[T] {
	return &Calc[T]{d}
}
func (c *Calc[T]) Add(d T) *Calc[T] {
	c.data += d
	return c
}

func (c *Calc[T]) Sub(d T) *Calc[T] {
	c.data -= d
	return c
}

func (c *Calc[T]) Mul(d T) *Calc[T] {
	c.data *= d
	return c
}

func (c *Calc[T]) Div(d T) *Calc[T] {
	c.data /= d
	return c
}

func (c *Calc[T]) Get() T {
	return c.data
}

type INumber interface {
	int | uint | uint8 | int8 | uint16 | int16 | uint32 | int32 | uint64 | int64 | float32 | float64
}
