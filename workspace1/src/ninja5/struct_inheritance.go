package main

import "fmt"

type person struct{
	first string
	last string
	age int
}

type secretAgent struct{
	person
	ltk bool
}

func main() {
	s1 := secretAgent{
		person: person{
			first: "James",
			last: "Bond",
		},
		ltk: true,
	}

	s2 := secretAgent{
		person: person{
			first: "Penny",
			last: "Menny",
			age: 31,
		},
		ltk: true,
	}

	fmt.Println(s1)
	fmt.Println(s1.first)
	fmt.Println(s2)
	fmt.Println(s2.last)//You can write s2.person.first in case of collisions.
}
