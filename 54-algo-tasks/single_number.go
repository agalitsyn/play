// Level: easy
//
// Source: https://leetcode.com/problems/single-number/
//
// Given an array of integers. All numbers occur twice except one number which occurs once. Find the number in O(n) time & constant extra space.
// Example :
// Input:  ar[] = {7, 3, 5, 4, 5, 3, 4}
// Output: 7
//
// One solution is to check every element if it appears once or not. Once an element with a single occurrence is found, return it. Time complexity of this solution is O(n2).
// A better solution is to use hashing.
// 1) Traverse all elements and put them in a hash table. Element is used as key and the count of occurrences is used as the value in the hash table.
// 2) Traverse the array again and print the element with count 1 in the hash table.
// This solution works in O(n) time but requires extra space.
//
// The best solution is to use XOR. XOR of all array elements gives us the number with a single occurrence. The idea is based on the following two facts.
// a) XOR of a number with itself is 0.
// b) XOR of a number with 0 is number itself.
//
// Let us consider the above example.
// Let ^ be xor operator as in C and C++.
// res = 7 ^ 3 ^ 5 ^ 4 ^ 5 ^ 3 ^ 4
//
// Since XOR is associative and commutative, above
// expression can be written as:
// res = 7 ^ (3 ^ 3) ^ (4 ^ 4) ^ (5 ^ 5)
//     = 7 ^ 0 ^ 0 ^ 0
//     = 7 ^ 0
//     = 7
package main

import "fmt"

func main() {
	fmt.Println(Solve([]int{9, 2, 3, 2, 3}))
}

func Solve(data []int) int {
	x := 0
	for _, a := range data {
		x ^= a
	}
	return x
}
