package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
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

	ch := make(chan string)
	quit := make(chan bool)
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := stdout.Read(buf)
			if n != 0 {
				ch <- string(buf[:n])
			}
			if err != nil {
				if err == io.EOF {
					log.Println("[DEBUG]: End of file")
					break
				}
				log.Fatal(err.Error())
			}
		}

		log.Println("[DEBUG]: Goroutine finished")

		close(ch)
	}()

	time.AfterFunc(time.Second, func() { quit <- true })

	for {
		select {
		case s, ok := <-ch:
			if !ok {
				os.Exit(0)
			}
			fmt.Print(s)
		case <-quit:
			cmd.Process.Kill()
		}
	}

}
