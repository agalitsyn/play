// https://tour.golang.org/moretypes/20
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	strMap := make(map[string]int)
	strArray := strings.Fields(s)

	for i := 0; i < len(strArray); i++ {
		strMap[strArray[i]]++
	}

	return strMap
}

func main() {
	wc.Test(WordCount)
}