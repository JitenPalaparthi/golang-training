package main

func main() {

	// rows := db.query("users").where("status=active").limit(10).offset(10).exec() //chain of actions, it is just example but not really similar query there in Go

	c := New(10).Add(10.10).Mul(2.2).Add(int8(10)).Sub(10).Div(2).Mul(4).Div(1).Add(true).Mul(true).Add("Hello World").Get()
	println(c)

}

type Calc struct {
	data float64
}

func New(d float64) *Calc {
	return &Calc{d}
}
func (c *Calc) Add(d any) *Calc {
	if isNumber(d) {
		c.data += ToFloat64(d)
	}
	return c
}

func (c *Calc) Sub(d any) *Calc {
	if isNumber(d) {
		c.data -= ToFloat64(d)
	}
	return c
}

func (c *Calc) Mul(d any) *Calc {
	if isNumber(d) {
		c.data *= ToFloat64(d)
	}
	return c
}

func (c *Calc) Div(d any) *Calc {
	if isNumber(d) {
		c.data /= ToFloat64(d)
	}
	return c
}

func (c *Calc) Get() float64 {
	return c.data
}

func isNumber(n any) bool {
	switch n.(type) { // n.(type) gives the type of the variable
	case uint, int, uint8, int8, uint16, int16, uint32, int32, uint64, int64, float32, float64:
		return true
	default:
		return false
	}
}

func ToFloat64(b any) float64 {
	switch b.(type) {
	case uint8:
		return float64(b.(uint8))
	case uint16:
		return float64(b.(uint16))
	case uint32:
		return float64(b.(uint32))
	case uint64:
		return float64(b.(uint64))
	case uint:
		return float64(b.(uint))
		// lets use v for few types
	case int8:
		return float64(b.(int8))
	case int16:
		return float64(b.(int16))
	case int32:
		return float64(b.(int32))
	case int64:
		return float64(b.(int64))
	case int:
		return float64(b.(int))
	case float32:
		return float64(b.(float32))
	case float64:
		return b.(float64)
	default:
		return 0 // the defailt is zero then multiplication and division has a problem
	}
}
