package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type process struct {
	command *exec.Cmd
	stream  io.Reader
}

func newProcess(cmd *exec.Cmd) *process {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err.Error())
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err.Error())
	}
	return &process{
		command: cmd,
		stream:  stdout,
	}
}

func (r *process) Read(p []byte) (n int, err error) {
	n, err = r.stream.Read(p)
	return n, err
}

func main() {
	var logFile string
	if len(os.Args) > 1 {
		logFile = os.Args[1]
	} else {
		log.Fatal("file name is required")
	}

	cmd := exec.Command("cat", logFile)
	process := newProcess(cmd)

	lines, err := lineCounter(process)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("lines = %+v\n", lines)
}

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)
		fmt.Println(string(buf[:c]))

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
