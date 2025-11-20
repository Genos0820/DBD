package main

import "fmt"

type geometery interface {
	area() float32
}

type rect struct {
	width  float32
	length float32
}

type circle struct {
	radius float32
}

func (r rect) area() float32 {
	return r.length * r.width
}

func (c circle) area() float32 {
	return 2 * 3.14 * c.radius
}

func measure(g geometery) {
	fmt.Printf("Type: %T/n", g)
	fmt.Println("Area:", g.area())
}

func main() {
	r := rect{
		width:  2.0,
		length: 3.0,
	}
	c := circle{
		radius: 3.0,
	}
	measure(r)
	measure(c)
}
