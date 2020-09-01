// Level: easy
//
// Source: https://leetcode.com/problems/maximum-subarray/
//
// Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

// Example:
// Input: [-2,1,-3,4,-1,2,1,-5,4],
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.

// Follow up:
// If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
package main

import (
	"fmt"
	"math"
)

// Kadane's algorithm
func Solve(nums []int) int {
	start, stop := 0, 0

	best_sum := math.MinInt64
	current_sum := 0
	for i := 0; i < len(nums); i++ {
		// 1. update current sum with new value of array
		current_sum += nums[i]

		// 2. update best sum if current cum is bigger
		if current_sum > best_sum {
			best_sum = current_sum

			stop = i
		}

		// 3. drop previous part of array if it's negative
		if current_sum < 0 {
			current_sum = 0

			start = i + 1
		}
	}
	fmt.Printf("Found Maximum Subarray between %d and %d\n", start, stop)
	return best_sum
}

func main() {
	in := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(Bruteforce(in))
	fmt.Println(Solve(in))
}

func Bruteforce(nums []int) int {
	start, stop := 0, 0

	sum := 0
	for left := 0; left < len(nums); left++ {
		partial_sum := 0
		for right := left; right < len(nums); right++ {
			partial_sum += nums[right]
			if partial_sum > sum {
				sum = partial_sum

				start = left
				stop = right
			}
		}
	}
	fmt.Printf("Found Maximum Subarray between %d and %d\n", start, stop)
	return sum
}
