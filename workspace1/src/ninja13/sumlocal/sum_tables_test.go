package main

import "testing"

func TestMySumTables(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{[]int{21, 21}, 42},
		{[]int{2, 2}, 4},
		{[]int{1, -1, 0}, 0},
	}

	for _, v := range tests {
		want := v.answer
		if got := mySumLocal(v.data...); got != want {
			t.Error("Expected", want, "Got", got)
		}
	}
}
