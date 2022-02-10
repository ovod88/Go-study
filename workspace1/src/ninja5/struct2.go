package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	first string
	last  string
	contactInfo
}

func main() {
	p1 := person{
		first: "Vova",
		last:  "Khrystenko",
		contactInfo: contactInfo{
			email: "vova@gmail.com",
			zip:   85869,
		},
	}
	p1.updateName2("Linas")
	fmt.Println(p1)
}

func (p person) updateName(newname string) {
	p.first = newname
}

func (p *person) updateName2(newname string) {
	p.first = newname
}
