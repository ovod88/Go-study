package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("This should not print1")
	case 2 == 4:
		fmt.Println("This should not print2")
	case 4 == 4:
		fmt.Println("This should print1")
		fallthrough
	case 5 == 5:
		fmt.Println("This should print2")
		fallthrough
	case 7 == 5:
		fmt.Println("This should not print3")
		fallthrough
	case 6 == 5:
		fmt.Println("This should not print4")
		fallthrough
	case true:
		fmt.Println("This should not print5")
	default:
		fmt.Println("Default case")
	}
	switch "Bond" {
	case "Mommy", "bond":
		fmt.Println("No James Bond")
	case "Bond":
		fmt.Println("This is James Bond")
	}
}
