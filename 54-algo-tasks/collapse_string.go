// Level: easy
//
// Source: interview
//
// Collapse similar symbols in the given string.
//
// Example:
// In: SSSXYZBBAAA
// Out: S3XYZB2A3
package main

import "strconv"

func Solve(data string) string {
	if len(data) < 1 {
		return data
	}

	res := ""
	count := 1
	for i := 0; i < len(data)-1; i++ {
		if data[i] == data[i+1] {
			count++
		} else {
			res += string(data[i])
			if count > 1 {
				res += strconv.Itoa(count)
			}
			count = 1
		}
	}

	res += string(data[len(data)-1])
	if count > 1 {
		res += strconv.Itoa(count)
	}

	if len(res) > len(data) {
		return data
	}
	return res
}
