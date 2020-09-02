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
		{
			[]int{0, 1, 0, 3, 12},
			[]int{1, 3, 12, 0, 0},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			Solve(tt.in)
			if !reflect.DeepEqual(tt.out, tt.in) {
				t.Errorf("got %+v, want %+v", tt.in, tt.out)
			}
		})
	}
}
