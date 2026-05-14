package cuboid

type Cuboid struct {
	L, B, H float32
}

func NewCuboid(l, b, h float32) *Cuboid {
	return &Cuboid{l, b, h}
}

func (r *Cuboid) Area() float64 { // receiver name can be any valid identifer , just to know that I kept it as r
	return float64(r.L) * float64(r.B) * float64(r.H)
}

func (c *Cuboid) Perimeter() float64 {
	return 2 * float64(c.L+c.B+c.H)
}
