package main

import (
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])

	if err != nil {
		os.Exit(1)
	} else {
		io.Copy(os.Stdout, f)
	}

}
