package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	data := make(chan int)
	ticker := time.Tick(100 * time.Millisecond)

	go func() {
		defer close(data)
		for i := 1; i <= 10; i++ {
			data <- i
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}
	}()

	for {
		select {
		case i, ok := <-data:
			if !ok {
				fmt.Printf("chan is closed, and returns zero-value of type int: %v\n", <-data)
				return
				// or
				//data = nil
				//continue
				// and change loop to
				// for data != nil {
			}
			fmt.Println(i)
		case <-ticker:
			fmt.Println("tick")
		}
	}
}
