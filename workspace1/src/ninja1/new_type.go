package main

import "fmt"

type hotdog int

var x hotdog
var y int = 43

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	fmt.Println(y)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	// z := 42
	// fmt.Println(z)
	// z = "example"
	// fmt.Println(z)
}
