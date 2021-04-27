package main

import "fmt"

func main() {
	fmt.Println(mySumLocal(20, 43, 22, 32))
}

func mySumLocal(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
