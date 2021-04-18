package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)

	c <- 42

	fmt.Println(<-c)

	c <- 43

	fmt.Println(<-c)
// 	fmt.Println(<-c)
}