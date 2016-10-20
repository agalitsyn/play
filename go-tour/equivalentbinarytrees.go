// https://tour.golang.org/concurrency/8
package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, results chan int) {
	var walker func(*tree.Tree)
	walker = func(t *tree.Tree) {
		if t != nil {
			walker(t.Left)
			results <- t.Value
			walker(t.Right)
		}
	}
	walker(t)

	close(results)
}

func Same(t1, t2 *tree.Tree) bool {
	t1Res, t2Res := make(chan int), make(chan int)
	go Walk(t1, t1Res)
	go Walk(t2, t2Res)

	for {
		t1El, t1Ended := <-t1Res
		t2El, t2Ended := <-t2Res
		if (t1El != t2El) || (t1Ended != t2Ended) {
			return false
		} else if t1Ended == t2Ended == true {
			return true
		}
	}

	return true
}

func main() {
	results := make(chan int)
	go Walk(tree.New(1), results)
	for r := range results {
		fmt.Println(r)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}