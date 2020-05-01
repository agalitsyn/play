// Level: easy
//
// Source: https://habr.com/ru/company/yandex/blog/449890/
// https://contest.yandex.ru/contest/8458/problems/?nc=kjmvs8aV
// https://rosettacode.org/wiki/Jewels_and_Stones#Go
//
// Create a function which takes two string parameters: 'stones' and 'jewels' and returns an integer.
// Both strings can contain any number of upper or lower case letters. However, in the case of 'jewels', all letters must be distinct.
// The function should count (and return) how many 'stones' are 'jewels' or, in other words, how many letters in 'stones' are also letters in 'jewels'.
//
// Example:
// Stones: "aAAbbbb"
// Jewels: "aA"
// 3
//
package main

import "fmt"

func main() {
	fmt.Println(js("aAAbbbb", "aA"))
}

func js(stones, jewels string) int {
	return 3
}
