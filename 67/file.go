package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logFile, err := os.Create("log")
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	for i := 0; i < 100; i++ {
		fmt.Fprintf(logFile, "log %d\n", i)
	}
}
