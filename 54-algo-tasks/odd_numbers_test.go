package main

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_Solve(t *testing.T) {
	var tests = []struct {
		in  []int
		out []int
	}{
		{[]int{2, 3, 4, 5, 7, 1}, []int{2, 1, 4, 3, 5, 7}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{0}, []int{0}},
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
