package main

import "fmt"

func main() {
	m := map[string]int{
		"James":        32,
		"Monny Monnay": 18,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barbabas"])

	v, ok := m["Barbabas"]
	fmt.Println(v)
	fmt.Println(ok)

	// if _, ok := m["Barbabas"]; !ok {
	// 	fmt.Println("Values Barbabas doesnt exists in map")
	// }

	// m["Todd"] = 33

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// delete(m, "Todd")
	// fmt.Println("")

	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	m2 := map[string][]string{
		"James": []string{"white", "red"},
		"Monny": []string{"yellow", "green"},
	}

	// fmt.Println(m2)

	for k, v := range m2 {
		for _, va := range v {
			fmt.Println(k, va)
		}
	}

	// fmt.Println(m2)

	// m2 := map[int][]string{
	// 	1: []string{"white", "red"},
	// 	2: []string{"yellow", "green"},
	// }

	// fmt.Println(m2)
}
