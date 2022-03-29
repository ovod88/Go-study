package main

import (
	"fmt"
)

func Append(slice []int, items ...int) []int {
	for _, item := range items {
		slice = Extend3(slice, item)
	}
	return slice
}

func Append2(slice []int, items ...int) []int {
	n := len(slice)
	total := len(slice) + len(items)
	if total > cap(slice) {
		// Reallocate. Grow to 1.5 times the new size, so we can still grow.
		newSize := total*3/2 + 1
		newSlice := make([]int, total, newSize)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:total]
	copy(slice[n:], items)
	return slice
}

func Extend3(slice []int, element int) []int {
	n := len(slice)
	if n == cap(slice) {
		// Slice is full; must grow.
		// We double its size and add 1, so if the size is zero we still grow.
		newSlice := make([]int, len(slice), 2*len(slice)+1)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

func main() {
	// slice := make([]int, 0, 5)
	// for i := 0; i < 10; i++ {
	// 	slice = Append(slice, i)
	// 	fmt.Printf("len=%d cap=%d slice=%v\n", len(slice), cap(slice), slice)
	// 	fmt.Println("address of 0th element:", &slice[0])
	// }
	// slice := []int{0, 1, 2, 3, 4}
	// fmt.Println(slice)
	// fmt.Printf("len=%d cap=%d slice=%v\n", len(slice), cap(slice), slice)
	// fmt.Println("address of 0th element:", &slice[0])
	// slice = Append(slice, 5, 6, 7, 8)
	// fmt.Println(slice)
	// fmt.Printf("len=%d cap=%d slice=%v\n", len(slice), cap(slice), slice)
	// fmt.Println("address of 0th element:", &slice[0])
	slice1 := []int{0, 1, 2, 3, 4}
	slice2 := []int{55, 66, 77}
	fmt.Println(slice1)
	slice1 = Append(slice1, slice2...) // The '...' is essential!
	fmt.Println(slice1)
}
