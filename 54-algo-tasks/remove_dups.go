// Level: easy
//
// Source: https://habr.com/ru/company/yandex/blog/449890/
// https://contest.yandex.ru/contest/8458/problems/C/
//
// Remove duplicates from unsorted array of integers.
//
// Example:
// In: [2, 4, 8, 8, 8]
// Out: [2, 4, 8]
package main

func RemoveDups(data []int) []int {
	var res []int
	uniques := make(map[int]bool)
	for i := 0; i < len(data); i++ {
		if _, ok := uniques[data[i]]; !ok {
			uniques[data[i]] = true
			res = append(res, data[i])
		}
	}
	return res
}
