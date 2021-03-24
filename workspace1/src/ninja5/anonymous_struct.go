package main

import "fmt"

func main() {
	p1 := struct{
		first string
		last string
		age int
	}{
		first: "James",
		last: "Bond",
		age: 30,
	}
	fmt.Println(p1)
}
