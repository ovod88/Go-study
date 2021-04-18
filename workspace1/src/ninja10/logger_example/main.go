package main

import (
	"fmt"
	// "log"
	"logger_example/logger"
	"os"
	"os/signal"
	"time"
)

type device struct {
	problem bool
}

func (d *device) Write(p []byte) (n int, err error) {
	for d.problem {
		time.Sleep(time.Second)
	}

	fmt.Print(string(p))
	return len(p), nil
}

func main() {
	const grs = 10

	var d device
	// var l log.Logger
	l := logger.New(&d, grs)
	// l.SetOutput(&d)

	for i := 0; i < grs; i++ {
		go func(id int){
			for {
				l.Println(fmt.Sprintf("%d: log some data\n", id))
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}

	// fmt.Println("TEST")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	
	// fmt.Println(<-sigChan)

	for{
		// fmt.Println("TEST BEFORE")
		<-sigChan
		// fmt.Println("TEST AFTER")
		d.problem = !d.problem
	}
}