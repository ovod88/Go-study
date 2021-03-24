package main

import "fmt"

func sum(desc string, x ...int) int {
	var sum int

	fmt.Println(desc)
	for _, v := range x {
		fmt.Println(sum)
		sum += v
	} 
	return sum
}

func main() {
	xi := []int{1,2,3,4,5,6,7}
	fmt.Println(sum("", xi...)) 
	fmt.Println() 
	fmt.Println(sum("", 1,2,3,4,56,4)) 
	fmt.Println() 
	fmt.Println(sum(""))
}
