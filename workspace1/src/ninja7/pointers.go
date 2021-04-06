package main

import "fmt"

func main() {
	a := 42
	var b *int
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b = &a
	fmt.Println(b)
	fmt.Println(*b)

	*b = 43
	fmt.Println(a)
}