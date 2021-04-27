package main

import "testing"

func TestMySum(t *testing.T) {
	want := 5
	if got := mySumLocal(2, 3); got != want {
		t.Error("Expected", want, "Got", got)
	}
}
