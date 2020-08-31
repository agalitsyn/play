// Level: easy
//
// Source: https://leetcode.com/problems/reverse-integer/
//
// Given a 32-bit signed integer, reverse digits of an integer.
//
// Example 1:
// Input: 123
// Output: 321
//
// Example 2:
// Input: -123
// Output: -321
//
// Example 3:
// Input: 120
// Output: 21
//
// Note:
// Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
package main

func Solve(x int) int {
	const MaxInt32 int = 1<<31 - 1
	const MinInt32 int = -1 << 31
	var rev int
	for x != 0 {
		if rev > MaxInt32/10 || rev < MinInt32/10 {
			return 0
		}

		pop := x % 10
		x /= 10
		rev = rev*10 + pop
	}
	return rev
}
