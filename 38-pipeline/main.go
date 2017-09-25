package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rawData, readerErrs := readerProducer(bytes.NewReader([]byte("foo bar baz")))
	rawData2, readerErrs2 := readerProducer(badReader{})
	rawData3 := randomProducer()

	go func() {
		for {
			select {
			case err := <-readerErrs:
				if err != nil {
					log.Error(err)
				}
			case err := <-readerErrs2:
				if err != nil {
					log.Error(err)
				}
			}
		}
	}()

	processedData := upperCaseProcessor(ctx, rawData)
	processedData2 := upperCaseProcessor(ctx, rawData2)

	output := merge(ctx, processedData, processedData2, rawData3)

	signalChan := make(chan os.Signal, 1)
	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		s := <-signalChan
		log.Printf("captured %v, exiting", s)
		cancel()
	}()

	for data := range output {
		fmt.Println(string(data))
	}
	log.Debug("canceled")
}

func upperCaseProcessor(ctx context.Context, in <-chan []byte) <-chan []byte {
	out := make(chan []byte)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- bytes.ToUpper(n):
			case <-ctx.Done():
				log.Debug("cancel upperCaseProcessor")
				return
			}
		}
	}()
	return out
}

func merge(ctx context.Context, inputs ...<-chan []byte) <-chan []byte {
	var wg sync.WaitGroup
	out := make(chan []byte)

	output := func(c <-chan []byte) {
		defer wg.Done()
		for d := range c {
			select {
			case out <- d:
			case <-ctx.Done():
				log.Debug("cancel merge")
				return
			}
		}
	}

	wg.Add(len(inputs))
	for _, input := range inputs {
		go output(input)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

type badReader struct{}

func (badReader) Read([]byte) (int, error) {
	ri := rand.Intn(5)
	if ri == 1 {
		return 0, fmt.Errorf("bad reader")
	}
	if ri == 4 {
		return 0, io.EOF
	}
	return ri, nil
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandBytesString(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return b
}

func randomProducer() <-chan []byte {
	out := make(chan []byte)
	go func() {
		defer close(out)
		for {
			out <- RandBytesString(8)
		}
	}()
	return out
}

func readerProducer(r io.Reader) (chan []byte, chan error) {
	out := make(chan []byte)
	errs := make(chan error)
	go func() {
		defer close(out)
		defer close(errs)

		buf := make([]byte, 10)
		for {
			n, err := r.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				errs <- errors.Wrap(err, "could not read")
				break
			}
			out <- buf[0:n]
		}
	}()
	return out, errs
}
