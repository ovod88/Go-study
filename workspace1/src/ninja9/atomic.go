package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Routines:", runtime.NumGoroutine())
	var counter int64

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func(){
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter ", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Routines in the progress:", runtime.NumGoroutine())
	} 
	wg.Wait( )
	fmt.Println("Routines in the end:", runtime.NumGoroutine())
	fmt.Println("Counter -> ", counter)
}