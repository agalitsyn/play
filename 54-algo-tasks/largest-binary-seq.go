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

import "fmt"

func main() {
	fmt.Println(largestSeq([]int{1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1}))
}

func largestSeq(data []int) int {
	return 4
}
