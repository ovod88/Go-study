package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	for i := 0; i < 5; i++ {
		fmt.Println("foo called")
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 5; i++ {
		fmt.Println("bar called")
	}
}

func main() {

	wg.Add(1)
	go foo() //starts goroutine
	bar()

	wg.Wait()
}