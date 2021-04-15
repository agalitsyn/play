package main

import "testing"

func BenchmarkOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		One()
	}
}

func BenchmarkTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Two()
	}
}

func TestOne(t *testing.T) {
	One()
}

func TestTwo(t *testing.T) {
	Two()
}
