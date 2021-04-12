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
	js := []byte(`[{"First":"Vova","Last":"Khrystenko","Age":32},{"First":"Lina","Last":"Khrystenko","Age":31}]`)

	people := []person{}
	if err := json.Unmarshal(js, &people); err == nil {
		fmt.Println(people)
	} else {
		fmt.Println(err)
	}

	for _, person := range people {
		fmt.Println(person.First)
	}
}