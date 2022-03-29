package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	const MN = 5
	args := os.Args[1:]
	var (
		ns [MN]int
		s  int
	)

	if len(args) == 0 {
		fmt.Println("Please provide numbers")
		return
	}

	for i := 0; i < MN; i++ {
		n, err := strconv.Atoi(args[i])
		if err == nil {
			fmt.Printf("Given number ->%v\n", n)
			ns[i] = n
			s += n
		}
	}
	fmt.Println(len(ns))
	fmt.Printf("Average is %v", s/len(ns))
}
