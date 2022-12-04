package main

type Rectangle interface {
	Perimeter() float64
	Area() float64
}

type Calculator struct {
	length float64
	width  float64
}

func (c Calculator) Perimeter() float64 {
	return 2 * (c.length + c.width)
}

func (c Calculator) Area() float64 {
	return c.length * c.width
}

func main() {
	var r Rectangle
	r = Calculator{20, 10}
	println(r.Perimeter())
	println(r.Area())
}
