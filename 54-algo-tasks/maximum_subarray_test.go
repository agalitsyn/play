package main

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_Solve(t *testing.T) {
	var tests = []struct {
		in  []int
		out int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{-2, -3, -1, -5}, -1},
		{[]int{0, 0}, 0},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			res := Solve(tt.in)
			if !reflect.DeepEqual(res, tt.out) {
				t.Errorf("got %+v, want %+v", res, tt.out)
			}
		})
	}
}
