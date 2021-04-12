package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last string
	Age int
}

func main() {
	p1 := person{
		First: "Vova",
		Last: "Khrystenko",
		Age: 32,
	}
	p2 := person{
		First: "Lina",
		Last: "Khrystenko",
		Age: 31,
	}

	people := []person{p1, p2,}

	if bs, err := json.Marshal(people); err == nil {
		fmt.Println(string(bs))
	} else {
		fmt.Println(err)
	}
}