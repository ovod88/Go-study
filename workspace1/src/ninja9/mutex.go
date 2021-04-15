package main

import (
	"fmt"
	"runtime"
	"sync"
	// "time"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Routines:", runtime.NumGoroutine())
	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func(){
			mu.Lock()
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Routines in the progress:", runtime.NumGoroutine())
	} 
	wg.Wait( )
	fmt.Println("Routines in the end:", runtime.NumGoroutine())
	fmt.Println("Counter -> ", counter)
}