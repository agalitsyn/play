// Level: easy
//
// Source: real interview
//
// Find median of unsorted array
//
// Example
// [2, 8, 5, 1, 4]
// 4
//
// [2, 8, 5, 1, 4, 5]
// 4.5
//
// [1, 1]
// 1.0
//
package main

import (
	"fmt"
	"sort"
)

func Solve(data []int) float64 {
	if len(data) < 1 {
		return 0.0
	}

	sort.Ints(data)
	fmt.Printf("%+v\n", data)

	if len(data)%2 == 0 {
		middle := len(data) / 2
		x, y := data[middle-1], data[middle]
		return (float64(x) + float64(y)) / 2.0
	}
	return float64(data[len(data)/2])
}
