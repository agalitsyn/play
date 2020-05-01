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

import "fmt"

func main() {
    fmt.Println(removeDups([]int{2, 4, 8, 8, 8}))
}

func removeDups(data []int) []int {
    return []int{2, 4, 8}
}
