package main

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_Anagram(t *testing.T) {
	var tests = []struct {
		w1 string
		w2 string
		out bool
	}{
		{"qiu", "uiq", true},
		{"foo", "bar", false},
		{"", "bar", false},
		{"foo", "", false},
		{"", "", false},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			res := Anagram(tt.w1, tt.w2)
			if !reflect.DeepEqual(res, tt.out) {
				t.Errorf("got %+v, want %+v", res, tt.out)
			}
		})
	}
}
