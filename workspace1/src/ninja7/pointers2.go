package main

import "fmt"

type person struct {
	firstname string
	lastname string
}

func changeme(p *person) {
	(*p).firstname = "Vova"
	p.lastname = "Khrystenko"//The same as (*p).lastname == p.lastname
}

func main() {
	p1 := person{
		firstname: "James",
		lastname: "Bond",
	}
	fmt.Println(p1.firstname, p1.lastname)
	changeme(&p1)
	fmt.Println(p1.firstname, p1.lastname)
}