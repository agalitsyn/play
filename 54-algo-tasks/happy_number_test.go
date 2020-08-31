package main

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_Solve(t *testing.T) {
	var tests = []struct {
		in  int
		out bool
	}{
		{19, true},
		{1945, false},
		{0, false},
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
