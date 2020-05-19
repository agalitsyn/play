// Level: easy
//
// Source: interview
//
// Collapse similar symbols in the given string.
//
// Example:
// In: SSSXYZBBAAA
// Out: S3XYZB2A2
package main

import (
	"fmt"
	"strconv"
)

func CollapseString(data string) string {
	res := ""
	var prev rune
	count := 0

	for i, s := range data {
		fmt.Printf("%s\n", string(s))
		fmt.Printf("%v\n", i)
		if prev != s {
			if count != 0 {
				res += string(strconv.Itoa(count + 1))
			}

			res += string(s)
			count = 0
			prev = s
		} else {
			count++
		}
	}
	return res
}
