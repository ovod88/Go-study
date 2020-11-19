package main

import "fmt"

func main() {
	x := []int{32, 49, 22} //Slice
	x[1] = 11
	fmt.Println(x[:1])

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
