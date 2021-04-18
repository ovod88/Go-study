package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 2)

	go func() {
		for{
			fmt.Println("Waiting for output from channel in goroutine")
			fmt.Println(<- c)
			fmt.Println("Done waiting for output from channel in goroutine")
		}
	}()
	time.Sleep(3 * time.Second)

	c <- 43

	// fmt.Println("Added value 43 to channel")

	time.Sleep(3 * time.Second)

	c <- 44

	time.Sleep(3 * time.Second)
	// fmt.Println("Added value 44 to channel")	
	c <- 45

	time.Sleep(3 * time.Second)
}