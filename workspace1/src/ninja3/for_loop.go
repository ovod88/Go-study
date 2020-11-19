package main

import "fmt"

func main() {
	// x := 1
	y := 1
	// for x < 10 {
	// fmt.Println(x)
	// x++
	// }

	for {
		if y > 10 {
			break
		}
		fmt.Println(y)
		y++
	}
}
