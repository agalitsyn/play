// Level: easy
//
// Source: https://habr.com/ru/company/yandex/blog/449890/
// https://contest.yandex.ru/contest/8458/problems/B/
//
// Find the largest sequence of 1s in binary vector.
//
// Example:
// [1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1]
// 4
package main

func Solve(data []int) int {
	if len(data) < 1 {
		return 0
	}

	count := 1
	prev := 0
	for i := 0; i < len(data); i++ {
		if prev == data[i] {
			count++
		}
	}
	return count
}
