package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	logFile := os.Args[1]
	cmd := exec.Command("tail", "-n", "+1", "-f", logFile)
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
