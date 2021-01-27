package morestrings

import (
	"fmt"
	"testing"
)

func TestReverseRunes(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	fmt.Println(cases)
	for _, c := range cases {
		if got := ReverseRunes(c.in); got != c.want {
			t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
