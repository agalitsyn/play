package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatal("missing command to execute")
	}
	cmdName := flag.Arg(0)
	cmdArgs := flag.Args()[1:]

	out := recieve(os.Stdin)
	for b := range out {
		// no reuse :( https://golang.org/pkg/os/exec/#Cmd
		cmd := exec.Command(cmdName, cmdArgs...)

		const placeholder = `{}`
		for i := range cmd.Args {
			if cmd.Args[i] == placeholder {
				cmd.Args[i] = string(b)
				break
			}
		}

		// user might want to pass commands like `time echo 1`, which logs to stderr
		cmdOut, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("could not execute command: %v", err)
		}

		fmt.Print(string(cmdOut))
	}
}

func recieve(r io.Reader) <-chan []byte {
	out := make(chan []byte)
	go func() {
		defer close(out)

		s := bufio.NewScanner(r)
		for s.Scan() {
			b := make([]byte, len(s.Bytes()))
			copy(b, s.Bytes())

			out <- b
		}

		if err := s.Err(); err != nil {
			log.Printf("could not read from pipe: %v", err)
		}
	}()
	return out
}
