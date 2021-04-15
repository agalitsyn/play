package main

import (
	"sync"
)

func One() {
	const parts = 3

	c := make(chan int, parts)
	go func() {
		defer close(c)

		var wg sync.WaitGroup
		for i := 0; i < parts; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				c <- i
			}(i)
		}
		wg.Wait()
	}()

	for res := range c {
		_ = res
	}
}

func Two() {
	const parts = 3

	c := make(chan int, parts)
	var wg sync.WaitGroup
	wg.Add(parts)

	go func() {
		defer close(c)
		wg.Wait()
	}()

	for i := 0; i < parts; i++ {
		go func(i int) {
			defer wg.Done()
			c <- i
		}(i)
	}

	for res := range c {
		_ = res
	}
}
