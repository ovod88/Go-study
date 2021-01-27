package main

import "fmt"

func main() {
	x := 2
	if x := 42; x != 2 {
		fmt.Println(x)
		fmt.Println("001")
	}
	fmt.Println(x)
}
