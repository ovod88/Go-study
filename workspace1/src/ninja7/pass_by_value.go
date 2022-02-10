package main

import "fmt"

func main() {
	// x := 0
	// fmt.Println(x)
	// foo(x)
	// fmt.Println(x)

	z := 0
	fmt.Println(z)
	foo2(&z)
	fmt.Println(z)
}

// func foo(y int) {
// 	y = 43
// 	fmt.Println(y)
// }

func foo2(y *int) {
	*y = 43
}
