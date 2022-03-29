package main

import (
	"fmt"
	"strings"
)

func main() {
	names := [...][3]string{
		{"First Name", "Last Name", "Nickname"},
		{"Albert", "Einstein", "emc2"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}

	for i := range names {
		n := names[i]
		fmt.Printf("%-15s %-15s %-15s\n", n[0], n[1], n[2])

		if i == 0 {
			fmt.Println(strings.Repeat("=", 50))
		}
	}

	for i, name := range names {
		fmt.Printf("%-15s %-15s %-15s\n", name[0], name[1], name[2])

		if i == 0 {
			fmt.Println(strings.Repeat("=", 50))
		}
	}
}
