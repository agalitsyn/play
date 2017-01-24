package main

import "fmt"

func main() {
	c := make(chan string)
	go func() {
		c <- "qwe"
	}()
	fmt.Println(<-c)
}
