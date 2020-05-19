package main

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_RemoveDups(t *testing.T) {
	var tests = []struct {
		in  []int
		out []int
	}{
		{[]int{2, 4, 8, 8, 8}, []int{2, 4, 8}},
		{[]int{1}, []int{1}},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			res := RemoveDups(tt.in)
			if !reflect.DeepEqual(res, tt.out) {
				t.Errorf("got %+v, want %+v", res, tt.out)
			}
		})
	}
}
