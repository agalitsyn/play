// Level: easy
// Tags: strings
//
// Source: https://habr.com/ru/company/yandex/blog/449890/
// https://contest.yandex.ru/contest/8458/problems/E/
//
// Find if string is an anagram.
//
// Example:
// In: qiu, uiq
// Out: true
package main

func Solve(w1, w2 string) bool {
	if w1 == "" || w2 == "" {
		return false
	}

	n := len(w1)
	runes := make([]rune, n)
	for _, rune := range w1 {
		n--
		runes[n] = rune
	}

	w2Rev := string(runes[n:])
	return w1 == w2Rev
}
