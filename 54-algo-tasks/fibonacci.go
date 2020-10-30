package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	a, b := 0, 1
	for i := 0; i < 10; i++ {
		fmt.Println(a + b)
		a, b = b, a+b
	}
}
