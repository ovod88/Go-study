package main

import "fmt"

func main() {
	x := []int{32} //Slice
	// x[1] = 11 //Throws an error
	fmt.Println(x)
	fmt.Printf("%T", x)
}
