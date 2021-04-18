package main

import (
	"fmt"
)

func send(e, o, q chan<- int) {
	for i:= 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Received value from even channel: ", v)
		case v := <-o:
			fmt.Println("Received value from odd channel: ", v)
		case v := <-q:
			fmt.Println("Received value from quit channel: ", v)
			return
		}
	}
}

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// c <- 42//blocks the code

	go send(even, odd, quit)

	receive(even, odd, quit)
}