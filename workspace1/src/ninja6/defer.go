package main

import "fmt"

func foo() {
	fmt.Println("hello from foo")
}

func foo2() {
	fmt.Println("hello from foo2")
}

func foo3() {
	fmt.Println("hello from foo3")
}

func bar() {
	fmt.Println("hello from bar")
}

func main() {
	defer foo()
	defer foo2()
	defer foo3()
	bar()
}
