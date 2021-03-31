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
	desc string
}

type shape interface {
	area() float32
}

func (s square) area() float32 {
	return s.a * s.b
}

func (c circle) area() float32 {//func (c *circle) area() float32 will work the same way
	return c.r * c.r * math.Pi
}

func info(s shape) {
	fmt.Println(s.area())
	// shape.desc = "Hello" - This womtwork because you are trying to change the value, it is ot possible in Go
}

func main() {
	square1 := square{
		a: 4,
		b: 5,
	}
	circle1 := circle{r:5,}

	info(square1)
	info(circle1)
	info(&circle1)//Works the same
}
