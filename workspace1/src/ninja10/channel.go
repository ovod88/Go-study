package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	// c <- 42//blocks the code

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}