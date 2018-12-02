package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type config struct {
	Rate       float64
	RatePeriod time.Duration
	Inflight   uint
}

func main() {
	var cfg config
	flag.Float64Var(&cfg.Rate, "rate", 0, "Limit of actions per one period for one parallel run.")
	flag.DurationVar(&cfg.RatePeriod, "rate-period", time.Second, "Period for rate actions.")
	flag.UintVar(&cfg.Inflight, "inflight", 1, "Limit for parallel runs.")

	var timer = flag.Bool("timer", false, "Run intergated timer")
	var debug = flag.Bool("debug", false, "Output debug info")

	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("missing command to execute")
	}

	if *debug {
		log.Printf("started with config: %+v", cfg)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigquit := make(chan os.Signal, 1)
	signal.Notify(sigquit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		s := <-sigquit
		log.Printf("captured %v, exiting...", s)
		cancel()
	}()

	if *timer {
		defer timeTaken(time.Now())
	}

	cmdName := flag.Arg(0)
	args := flag.Args()[1:]

	cmdArgs := make([]string, len(args))
	copy(cmdArgs, args)
	placeholderIndex, shouldReplace := indexPlaceholder(`{}`, cmdArgs)

	subProcess := func(replace string) {
		if shouldReplace {
			cmdArgs[placeholderIndex] = replace
		}

		if *debug {
			log.Printf("%s %v", cmdName, cmdArgs)
		}

		out, err := runCmd(cmdName, cmdArgs...)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Fprint(os.Stdout, out)
	}

	// no rate limiting, run and exit
	if cfg.Rate == 0 {
		for s := range readPipe(ctx, os.Stdin) {
			subProcess(s)
		}
		return
	}

	// with rate-limiter
	period := (float64(cfg.RatePeriod) / cfg.Rate) / float64(cfg.Inflight)
	limiter := time.Tick(time.Duration(period))
	queue := readPipe(ctx, os.Stdin)

	// spawn workers
	var wg sync.WaitGroup
	var i uint
	for i = 0; i < cfg.Inflight; i++ {
		wg.Add(1)
		go func(queue <-chan string, id uint) {
			defer wg.Done()

			for s := range queue {
				if *debug {
					log.Printf("worker %d got %s", id, s)
				}

				select {
				case <-ctx.Done():
					return
				case <-limiter:
					subProcess(s)
				}
			}
		}(queue, i)
	}
	if *debug {
		log.Printf("period: %v", period)
		log.Printf("%v workers", i)
	}

	wg.Wait()
}

func readPipe(ctx context.Context, r io.Reader) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)

		s := bufio.NewScanner(r)
		for s.Scan() {
			b := make([]byte, len(s.Bytes()))
			copy(b, s.Bytes())

			select {
			case out <- string(b):
			case <-ctx.Done():
				return
			}
		}

		if err := s.Err(); err != nil {
			log.Printf("could not read from pipe: %v", err)
		}
	}()
	return out
}

func runCmd(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	cmdOut, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("could not execute command: %v", err)
	}
	return string(cmdOut), nil
}

func indexPlaceholder(placeholder string, args []string) (int, bool) {
	for i := range args {
		if args[i] == placeholder {
			return i, true
		}
	}
	return 0, false
}

func timeTaken(t time.Time) {
	log.Printf("took %s\n", time.Since(t))
}
