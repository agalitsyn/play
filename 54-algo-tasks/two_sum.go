// Level: easy
//
// Source: https://leetcode.com/problems/two-sum/
//
// Given an array of integers nums and and integer target, return the indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

// Example 1:
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Output: Because nums[0] + nums[1] == 9, we return [0, 1]

// Example 2:
// Input: nums = [3,2,4], target = 6
// Output: [1,2]

// Example 3:
// Input: nums = [3,3], target = 6
// Output: [0,1]
package main

func Solve(nums []int, target int) []int {
	storage := make(map[int]int)
	for i, v := range nums {
		next := target - v
		if j, ok := storage[next]; ok {
			return []int{i, j}
		} else {
			storage[v] = i
		}
	}

	return []int{}
}
