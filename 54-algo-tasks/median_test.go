package main

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_Median(t *testing.T) {
	var tests = []struct {
		in  []int
		out float64
	}{
		{[]int{2, 8, 5, 1, 4}, 4.0},
		{[]int{2, 8, 5, 1, 4, 5}, 4.5},
		{[]int{1, 1, 1}, 1.0},
		{[]int{1, 1}, 1.0},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			res := Median(tt.in)
			if !reflect.DeepEqual(res, tt.out) {
				t.Errorf("got %+v, want %+v", res, tt.out)
			}
		})
	}
}
