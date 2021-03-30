package main

import (
	"fmt"
	"math"
)

type square struct {
	a float32
	b float32
}

type circle struct {
	r float32
}

type shape interface {
	area() float32
}

func (s square) area() float32 {
	return s.a * s.b
}

func (c circle) area() float32 {
	return c.r * c.r * math.Pi
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	square1 := square{
		a: 4,
		b: 5,
	}
	circle1 := circle{r:5,}

	info(square1)
	info(circle1)
}
