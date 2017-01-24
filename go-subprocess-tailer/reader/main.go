package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	Byte     = 1.0
	Kilobyte = 1024 * Byte
	Megabyte = 1024 * Kilobyte
)

// tailMaxDepth defines how many last lines will tail output with no filter set
const tailMaxDepth = 100

func main() {
	logFile := os.Args[1]

	f, err := os.Stat(logFile)
	if err != nil {
		log.Fatal(err)
	}

	lines := "+1"
	size := float32(f.Size())
	if size > Megabyte {
		lines = fmt.Sprintf("%v", tailMaxDepth)
	}

	cmd := exec.Command("tail", "--lines", lines, "--follow", logFile)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(stdout)
	s.Split(bufio.ScanLines)

	ch := make(chan string)
	go func() {
		for s.Scan() {
			line := s.Text()
			ch <- fmt.Sprintf("%v \n", line)
		}
		close(ch)
	}()

	for {
		select {
		case s, ok := <-ch:
			if !ok {
				os.Exit(0)
			}
			fmt.Print(s)
		}
	}
}
