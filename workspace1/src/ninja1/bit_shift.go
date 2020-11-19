package main

import "fmt"

const (
	a = iota
	b
	c
)

func main() {
	x := 2
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	y := x << 1
	fmt.Printf("%d\t\t%b\n", x, x)
	fmt.Printf("%d\t\t%b\n", y, y)
}
