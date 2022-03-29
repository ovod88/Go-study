package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	const (
		usage = "Please provide the amount to be converted."
		EUR   = iota
		GBP
		JPY
	)

	exrate := [...]float64{
		EUR: 0.88,
		GBP: 0.78,
		JPY: 113.02,
	}

	if len(os.Args) != 2 {
		fmt.Println(usage)
		return
	}

	dollars, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println(usage)
		return
	}
	fmt.Printf("%.2f USD is %.2f EUR\n", dollars, exrate[EUR]*dollars)
	fmt.Printf("%.2f USD is %.2f GBP\n", dollars, exrate[GBP]*dollars)
	fmt.Printf("%.2f USD is %.2f JPY\n", dollars, exrate[JPY]*dollars)
}
