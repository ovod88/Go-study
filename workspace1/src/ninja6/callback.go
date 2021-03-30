package main

import "fmt"

func sum(in ...int) int {
	var sum int

	for _, v := range in {
		// fmt.Println(sum)
		sum += v
	} 
	return sum
}

func even(f func(in ...int) int, in ...int) int {
	var result []int

	for _, v := range in {
		if v % 2 == 0 {
			result = append(result, v)
		} 
	}

	return f(result...)
}

func main() {
	fmt.Println("Count me please only even numbers")
	fmt.Println(even(sum, []int{1,2,3,4,5,6,7,8,67,5,5,4,4,4,5}...))
}