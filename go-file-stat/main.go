package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Stat(os.Args[1])
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("f = %+v\n", f.Size())
}
