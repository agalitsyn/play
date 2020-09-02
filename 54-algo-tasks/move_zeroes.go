// Level: easy
//
// Source: https://leetcode.com/problems/move-zeroes/
//
// Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.
//
// Example:
// Input: [0,1,0,3,12]
// Output: [1,3,12,0,0]
//
// Note:
// You must do this in-place without making a copy of the array.
// Minimize the total number of operations.
package main

import "fmt"

func Solve(nums []int) {
	nonZeroPointer := 0
	for _, x := range nums {
		if x != 0 {
			nums[nonZeroPointer] = x
			nonZeroPointer++
		}
	}
	for i := nonZeroPointer; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	Solve(nums)
	fmt.Println(nums)
}
