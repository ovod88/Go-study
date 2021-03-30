package main

import "fmt"

var x int

func main() {
	x++
	fmt.Println(x)
	foo()
	fmt.Println(x)

	{
		y := 42
		fmt.Println(y)
	}

	increment := incrementor()
	fmt.Println(increment())
}

func foo() {
	x++
}

func incrementor() func() int {
	var x int

	return func() int {
		x++
		return x//closure
	}
}