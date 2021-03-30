package main

import "fmt"

func foo(s string) {
	fmt.Println("Hello", s) 
}

var x int = 7
var g func() = func(){}

func woo(s string) string {
	return fmt.Sprint("Hello ", s) 
}

func mouse(ft string, lt int) (string, bool) {
	return ft, lt == 2
}

func retfunc() func() int {//func retfunc which returns func which returns int

	return func() int {
			fmt.Println("I am a very important function")
			return 42
		}	
}

func getfunc(input func() int) {//callback
	fmt.Println("Used inputed function")
	fmt.Println(input())
}

func main() {
	foo("Bond")
	fmt.Println(woo("cata")) 
	fmt.Println(mouse("Vova", 23))

	fmt.Println("")

	func(){
		fmt.Println("Anonymous function called")
	}()//Anonymous function

	fmt.Println("")

	f1 := func() {
		fmt.Println("This is a func expression")
	}// Func expression
	f1()

	fmt.Println("")

	f2 := retfunc()
	fmt.Println(f2())

	fmt.Println("")

	fmt.Println(retfunc()())


	fmt.Println("")
	getfunc(retfunc())
}
