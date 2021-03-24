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

func (sa person) speak(mesg string) {
	fmt.Println(sa.first, sa.last, "says ", mesg)
}

func (sa secretAgent) stay(mesg string) {
	fmt.Println(sa.first, sa.last, "says ", mesg)
}

func main() {
	s1 := secretAgent{
		person: person{
			first: "James",
			last: "Bond",
		},
		ltk: true,
	}

	fmt.Println(s1)
	s1.speak("heya")
}
