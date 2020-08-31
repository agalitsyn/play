// Level: easy
//
// Source: interview
//
// Sort odd numbers in given array, even should save their position
//
// Example:
// In:  [2, 3, 4, 5, 7, 1]
// Out: [2, 1, 4, 3, 5, 7]
package main

import (
	"sort"
)

func Solve(data []int) []int {
	var oddNumbers []int
	for i := 0; i < len(data); i++ {
		if data[i]%2 != 0 {
			oddNumbers = append(oddNumbers, data[i])
		}
	}
	sort.Ints(oddNumbers)

	oInd := 0
	for i := 0; i < len(data); i++ {
		if data[i]%2 != 0 && oInd < len(oddNumbers) {
			data[i] = oddNumbers[oInd]
			oInd++
		}
	}

	return data
}
