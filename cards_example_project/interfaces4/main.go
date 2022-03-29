package main

import "fmt"

type Puller interface {
	getMessage() string
}

type person struct {
	name string
}

func (person) getMessage() string {
	return "Hello here"
}

func (p *person) changeName(n string) {
	p.name = n
}

func printmessage(p Puller) {
	fmt.Println(p.getMessage())
}

func main() {
	p := person{
		name: "vova",
	}

	printmessage(&p)
	(&p).changeName("lina")
	fmt.Println(p)
}
