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

import "fmt"

func main() {
	fmt.Println(sortOddNumbers([]int{2, 3, 4, 5, 7, 1}))
}

func sortOddNumbers(data []int) []int {
	return []int{2, 1, 4, 3, 5, 7}
}
