package main

func main() {

	// rows := db.query("users").where("status=active").limit(10).offset(10).exec() //chain of actions, it is just example but not really similar query there in Go

	c := New(10).Add(10).Mul(2).Add(10).Sub(10).Div(2).Mul(4).Div(1).Get()
	println(c)

}

type Calc struct {
	data float64
}

func New(d float64) *Calc {
	return &Calc{d}
}
func (c *Calc) Add(d float64) *Calc {
	c.data += d
	return c
}

func (c *Calc) Sub(d float64) *Calc {
	c.data -= d
	return c
}

func (c *Calc) Mul(d float64) *Calc {
	c.data *= d
	return c
}

func (c *Calc) Div(d float64) *Calc {
	c.data /= d
	return c
}

func (c *Calc) Get() float64 {
	return c.data
}
