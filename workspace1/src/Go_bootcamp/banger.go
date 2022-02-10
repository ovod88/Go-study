package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]
	l := len(msg)

	surroundings := strings.Repeat("!", l)
	s := surroundings + strings.ToUpper(msg) + surroundings

	fmt.Println(s)
}
