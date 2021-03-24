package main

import "fmt"

type person struct{
	first string
	last string
}

func main() {
	p1 := person{
		first: "Vova",
		last: "Khrystenko",
	}
	fmt.Println(p1)
	fmt.Println(p1.first)
}
