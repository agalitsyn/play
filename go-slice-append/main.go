package main

import "fmt"

func main() {
	var array []string

	toAdd := []string{"foo", "bar", "baz"}
	for _, i := range toAdd {
		if i == "foo" {
			continue
		}
		array = append(array, i)
	}

	fmt.Println(array)
}
