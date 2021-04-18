package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)

	c <- 42
	// c <- 43//wont run until read

	fmt.Println(<-c)
}