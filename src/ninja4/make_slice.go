package main

import "fmt"

func main() {
	x := make([]int, 10, 20) //Slice with 10 filled elements and underlying 20 elements array
	x[9] = 42
	x = append(x, 2345)
	fmt.Println(x)
	x = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(x)
}
