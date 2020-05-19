package main

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_CollapseString(t *testing.T) {
	var tests = []struct {
		in  string
		out string
	}{
		{"SSSXYZBBAAA", "S3XYZB2A3"},
		{"ABC", "ABC"},
		{"", ""},
		{"AAAA", "A4"},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			res := CollapseString(tt.in)
			if !reflect.DeepEqual(res, tt.out) {
				t.Errorf("got %+v, want %+v", res, tt.out)
			}
		})
	}
}
