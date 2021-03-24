package main

import "fmt"

func foo(s string) {
	fmt.Println("Hello", s) 
}

func woo(s string) string {
	return fmt.Sprint("Hello ", s) 
}

func mouse(ft string, lt int) (string, bool) {
	return ft, lt == 2
}

func main() {
	foo("Bond")
	fmt.Println(woo("cata")) 
	fmt.Println(mouse("Vova", 23))
}
