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
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	ctx, cancel := context.WithCancel(context.Background())

	signalChan := make(chan os.Signal, 1)
	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	rawData, readerErrs := readerProducer(ctx, bytes.NewReader([]byte("foo bar baz")))
	rawData2, readerErrs2 := readerProducer(ctx, badReader{})
	rawData3 := randomProducer(ctx, 250*time.Millisecond)
	rawData4 := randomProducer(ctx, 750*time.Millisecond)

	errs := mergeErrors(ctx, readerErrs, readerErrs2)

	processedData := upperCaseProcessor(ctx, rawData)
	processedData2 := upperCaseProcessor(ctx, rawData2)
	processedData3 := upperCaseProcessor(ctx, rawData3)
	processedData4 := upperCaseProcessor(ctx, rawData4)

	output := merge(ctx, processedData, processedData2, processedData3, processedData4)

	go func() {
		s := <-signalChan
		log.Printf("captured %v, exiting", s)
		cancel()
	}()
	go func() {
		for err := range errs {
			log.Error(err)
		}
		log.Info("errors done")
	}()

	for data := range output {
		fmt.Println(string(data))
	}
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

func randomProducer(ctx context.Context, pause time.Duration) <-chan []byte {
	out := make(chan []byte)
	go func() {
		defer close(out)
		for {
			select {
			case out <- RandBytesString(8):
			case <-ctx.Done():
				log.Debug("cancel random producer")
				return
			}
			time.Sleep(pause)
		}
	}()
	return out
}

func readerProducer(ctx context.Context, r io.Reader) (chan []byte, chan error) {
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

			select {
			case out <- buf[0:n]:
			case <-ctx.Done():
				log.Debug("cancel reader producer")
				return
			}
		}
	}()
	return out, errs
}

func mergeErrors(ctx context.Context, inputs ...<-chan error) <-chan error {
	var wg sync.WaitGroup
	outCh := make(chan error)

	receiver := func(ch <-chan error) {
		defer wg.Done()
		for e := range ch {
			select {
			case outCh <- e:
			case <-ctx.Done():
				outCh <- errors.Wrap(ctx.Err(), "cancel merging errors")
				return
			}
		}
	}

	wg.Add(len(inputs))
	for _, input := range inputs {
		go receiver(input)
	}

	go func() {
		wg.Wait()
		close(outCh)
	}()

	return outCh
}
