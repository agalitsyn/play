package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/gravitational/trace"
)

func main() {
	var (
		filter  = flag.String("filter", "", "Filter applied to log file (supports grep like regexp)")
		limit   = flag.String("filter-limit", "10", "Limit filter output")
		logFile string
	)
	flag.Parse()
	if len(flag.Args()) > 0 {
		logFile = flag.Args()[0]
	} else {
		logFile = "."
	}

	// file could not exists yet, tail supports that case
	commands := []*exec.Cmd{}
	if *filter != "" {
		// grep asdf test.log | tail -n 3; tail -f -n0 test.log
		matcher := regexp.QuoteMeta(strings.TrimSpace(*filter))
		filterCmd := exec.Command("grep", "--line-buffered", "--extended-regexp", matcher, logFile)
		limitCmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("tail --lines %v; tail --follow --lines 0 %v", *limit, logFile))
		commands = append(commands, filterCmd, limitCmd)
	} else {
		readCmd := exec.Command("tail", "--lines", "100", "--follow", logFile, "--retry")
		commands = append(commands, readCmd)
	}

	pipeline, err := NewProcessGroup(commands...)
	log.Printf("tailing pipeline: %s", pipeline)
	if err != nil {
		log.Fatalf("failed to build a command pipeline: %v", err)
	}
	defer pipeline.Close()

	ch := OutputScanner(pipeline)
	for {
		select {
		case msg, ok := <-ch:
			if !ok {
				os.Exit(0)
			}
			fmt.Println(msg)
		}
	}
}

// OutputScanner spawns a goroutine to handle messages from the process group.
// Returns a channel where the received messages are sent to.
func OutputScanner(r io.Reader) chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		s.Split(bufio.ScanLines)
		for s.Scan() {
			ch <- string(s.Bytes())
		}
		log.Printf("closing tail message pump: %v", s.Err())
	}()
	return ch
}

// processGroup groups the processes that build a processing pipe
//
// implements fmt.Stringer
// implements io.Closer
// implements io.ReadCloser
// implements io.Reader
type processGroup struct {
	commands []*exec.Cmd
	closers  []io.Closer
	stream   io.Reader
}

func NewProcessGroup(commands ...*exec.Cmd) (group *processGroup, err error) {
	var stdout io.ReadCloser
	var closers []io.Closer
	for i, cmd := range commands {
		stdout, err = cmd.StdoutPipe()
		if err != nil {
			return nil, trace.Wrap(err)
		}
		closers = append(closers, stdout)
		cmd.Start()
		if i < len(commands)-1 {
			commands[i+1].Stdin = stdout
		}
	}

	return &processGroup{
		commands: commands,
		closers:  closers,
		stream:   stdout,
	}, nil
}

func (r *processGroup) Read(p []byte) (n int, err error) {
	n, err = r.stream.Read(p)
	return n, err
}

func (r *processGroup) Close() (err error) {
	// Close all open stdout handles
	for _, closer := range r.closers {
		closer.Close()
	}
	r.terminate()
	return trace.Wrap(err)
}

func (r *processGroup) String() string {
	var cmds []string
	for _, cmd := range r.commands {
		cmds = append(cmds, fmt.Sprintf("%v", cmd.Args))
	}
	return fmt.Sprintf("[%v]", strings.Join(cmds, ","))
}

// processTerminateTimeout defines the initial amount of time to wait for process to terminate
const processTerminateTimeout = 200 * time.Millisecond

func (r *processGroup) terminate() {
	terminated := make(chan struct{})
	head := r.commands[0]
	go func() {
		for _, cmd := range r.commands {
			// Await termination of all processes in the group to prevent zombie processes
			if err := cmd.Wait(); err != nil {
				log.Printf("%v exited with %v", cmd.Path, err)
			}
		}
		terminated <- struct{}{}
	}()

	if err := head.Process.Signal(syscall.SIGINT); err != nil {
		log.Printf("cannot terminate with SIGINT: %v", err)
	}

	select {
	case <-terminated:
		return
	case <-time.After(processTerminateTimeout):
	}

	if err := head.Process.Signal(syscall.SIGTERM); err != nil {
		log.Printf("cannot terminate with SIGTERM: %v", err)
	}

	select {
	case <-terminated:
		return
	case <-time.After(processTerminateTimeout * 2):
		head.Process.Kill()
	}
}
