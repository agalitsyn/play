package main

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_LargestSeq(t *testing.T) {
	var tests = []struct {
		in  []int
		out int
	}{
		{[]int{1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1}, 4},
		{[]int{1}, 1},
		{[]int{}, 0},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			res := LargestSeq(tt.in)
			if !reflect.DeepEqual(res, tt.out) {
				t.Errorf("got %+v, want %+v", res, tt.out)
			}
		})
	}
}
