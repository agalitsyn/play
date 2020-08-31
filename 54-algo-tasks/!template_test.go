package main

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_Solve(t *testing.T) {
	var tests = []struct {
		in  bool
		out bool
	}{
		{true, true},
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
