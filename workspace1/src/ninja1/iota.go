package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb
	gb
)

func main() {
	x := (42 == 42)
	fmt.Println(x)
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t\t%b\n", gb, gb)
}
