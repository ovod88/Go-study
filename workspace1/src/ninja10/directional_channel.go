package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan<- int, 2)//creates send-only channel
	b := make(<-chan int, 2)//creates receive-only channel

	c <- 43
	c <- 44
	b <- 1

	fmt.Println(<- c)

}