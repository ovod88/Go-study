package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

type human interface {
	speak(msg string)
}

type anybody interface{}

func (sa person) speak(mesg string) {
	fmt.Println(sa.first, sa.last, "speaks as person ", mesg)
}

// func (sa secretAgent) speak(mesg string) {
// 	fmt.Println(sa.first, sa.last, "speaks as sa ", mesg)
// }

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("It is a person", h.(person).first)
	case secretAgent:
		fmt.Println("It is a secret agent", h.(secretAgent).first)
	}
	fmt.Println("I called human in bar", h)
}

func all(smb anybody) { // the same as (smb interface{})
	fmt.Println("You are also somebody")
}

func personOnly(p person) {
	fmt.Println(p)
}

func main() {
	s1 := secretAgent{ // Here s1 is secretAgent and also human
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}

	// p1 := person{
	// 	first: "Dr.",
	// 	last:  "Yes",
	// }

	// fmt.Println(p1)
	bar(s1)
	// bar(p1)
	// all(s1)
	// personOnly(s1)
}
