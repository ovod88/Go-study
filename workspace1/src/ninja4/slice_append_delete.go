package main

import "fmt"

func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

func main() {
	x := []int{1, 2, 3, 4, 5} //Slice
	// x1 := [...]int{1, 2, 3, 4, 5} //Array
	// x2 := [3]int //Array
	// fmt.Printf("%T", x)
	fmt.Println(x)
	// x = append(x, 6, 7, 8, 9)
	// fmt.Println(x)
	// y := []int{20, 30, 40, 50}
	// x = append(x, y...)
	// fmt.Println(x)

	x = append(x[:2], x[4:]...)

	// array := [...]float64{7.0, 8.5, 9.1} //Array
	// x := Sum(&array)                     //C++ type behaviour
	fmt.Println(x)
}
