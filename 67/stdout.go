package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 100; i < 150; i++ {
		fmt.Fprintf(os.Stdout, "log %d\n", i)
	}
}
