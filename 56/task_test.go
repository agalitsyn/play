package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func Test_Merge2Channels(t *testing.T) {
	in1 := make(chan int)
	go genInts(in1)

	in2 := make(chan int)
	go genInts(in2)

	out := make(chan int)
	n := 100
	f := func(i int) int {
		r := rand.Intn(50)
		time.Sleep(time.Millisecond * time.Duration(r))

		return i + 1
	}

	Merge2Channels(f, in1, in2, out, n)

	expected := make([]int, n)
	for i := 0; i < n; i++ {
		expected[i] = (i + 1) + (i + 1)
	}

	actual := make([]int, 0, n)
	for i := range out {
		actual = append(actual, i)
	}
	sort.Ints(actual)

	for i, v := range expected {
		if actual[i] != v {
			t.Errorf("expected %v actual %v", v, actual[i])
		}
	}

}

func genInts(in chan<- int) {
	for i := 0; ; i++ {
		in <- i
	}
}
