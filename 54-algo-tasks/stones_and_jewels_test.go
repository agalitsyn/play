package main

import (
	"strconv"
	"testing"
)

func Test_StonesAndJewels(t *testing.T) {
	var tests = []struct {
		stones string
		jewels string
		out    int
	}{
		{"aAAbbbb", "aA", 3},
		{"ZZ", "z", 0},
		{"abc", "a", 1},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			res := StonesAndJewels(tt.stones, tt.jewels)
			if res != tt.out {
				t.Errorf("got %q, want %q", res, tt.out)
			}
		})
	}
}
