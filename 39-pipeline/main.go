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

	var wg sync.WaitGroup
	wg.Add(5)
	rawData1 := randomProducer(ctx, &wg)
	rawData2 := randomProducer(ctx, &wg)
	rawData3 := randomProducer(ctx, &wg)
	grc := &goodReadCloser{bytes.NewBufferString("foo bar baz")}
	brc := &badReadCloser{bytes.NewBufferString("foobar foobarbaz")}
	rawData4, readerErrs1 := readerProducer(ctx, &wg, grc)
	rawData5, readerErrs2 := readerProducer(ctx, &wg, brc)
	wg.Add(5)
	processedData1 := upperCaseProcessor(ctx, &wg, rawData1)
	processedData2 := upperCaseProcessor(ctx, &wg, rawData2)
	processedData3 := upperCaseProcessor(ctx, &wg, rawData3)
	processedData4 := upperCaseProcessor(ctx, &wg, rawData4)
	processedData5 := upperCaseProcessor(ctx, &wg, rawData5)
	wg.Add(2)
	output := merge(ctx, &wg, processedData1, processedData2, processedData3, processedData4, processedData5)
	errs := mergeErrors(ctx, &wg, readerErrs1, readerErrs2)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for err := range errs {
			if err != nil {
				log.Error(err)
			}
		}
	}()
	go func() {
		defer wg.Done()
		var i int
		for msg := range output {
			fmt.Printf("%d %s\n", i, string(msg))
			if i == 10 {
				// propagate cancel to all goroutines
				cancel()
				break
			}
			i++
		}
	}()

	wg.Add(1)
	signalChan := make(chan os.Signal, 1)
	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		defer wg.Done()
		s := <-signalChan
		log.Warnf("captured %v, exiting", s)
		cancel()
	}()

	wg.Wait()
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randBytesString(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return b
}

func randomProducer(ctx context.Context, wg *sync.WaitGroup) <-chan []byte {
	outCh := make(chan []byte)
	go func() {
		defer wg.Done()
		defer close(outCh)

		for {
			select {
			case outCh <- randBytesString(12):
			case <-ctx.Done():
				log.Debugf("random producer: %v", ctx.Err())
				return
			}
		}
	}()
	return outCh
}

func upperCaseProcessor(ctx context.Context, wg *sync.WaitGroup, in <-chan []byte) <-chan []byte {
	out := make(chan []byte)
	go func() {
		defer func() {
			if ctx.Err() != nil {
				log.Debugf("upper case processor: %v", ctx.Err())
			}
			close(out)
			wg.Done()
		}()

		for n := range in {
			select {
			case out <- bytes.ToUpper(n):
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func merge(ctx context.Context, wg *sync.WaitGroup, inputs ...<-chan []byte) <-chan []byte {
	var lwg sync.WaitGroup
	outCh := make(chan []byte)

	// Start an receiver goroutine for each input channel. receiver
	// copies values from ch to outCh until ch is closed, then calls wg.Done.
	receiver := func(ch <-chan []byte) {
		defer func() {
			if ctx.Err() != nil {
				log.Debugf("merge receiver: %v", ctx.Err())
			}
			lwg.Done()
		}()

		for i := range ch {
			select {
			case outCh <- i:
			case <-ctx.Done():
				return
			}
		}
	}

	lwg.Add(len(inputs))
	for _, input := range inputs {
		go receiver(input)
	}

	// Start a goroutine to close out once all the receiver goroutines are
	// done. This must start after the wg.Add call.
	go func() {
		defer wg.Done()

		lwg.Wait()
		close(outCh)
	}()

	return outCh
}

func mergeErrors(ctx context.Context, wg *sync.WaitGroup, inputs ...<-chan error) <-chan error {
	var lwg sync.WaitGroup
	outCh := make(chan error)

	receiver := func(ch <-chan error) {
		defer func() {
			if ctx.Err() != nil {
				log.Debugf("error merge receiver: %v", ctx.Err())
			}
			lwg.Done()
		}()

		for e := range ch {
			select {
			case outCh <- e:
			case <-ctx.Done():
				return
			}
		}
	}

	lwg.Add(len(inputs))
	for _, input := range inputs {
		go receiver(input)
	}

	go func() {
		defer wg.Done()

		lwg.Wait()
		close(outCh)
	}()

	return outCh
}

type badReadCloser struct {
	*bytes.Buffer
}

func (badReadCloser) Read([]byte) (int, error) {
	return 0, errors.Errorf("bad read")
}

func (badReadCloser) Close() error {
	return errors.Errorf("bad close")
}

type goodReadCloser struct {
	*bytes.Buffer
}

func (goodReadCloser) Close() error {
	return nil
}

func readerProducer(ctx context.Context, wg *sync.WaitGroup, r io.ReadCloser) (chan []byte, chan error) {
	out := make(chan []byte)
	errs := make(chan error)
	go func() {
		defer func() {
			if err := r.Close(); err != nil {
				errs <- errors.Wrap(err, "could not close")
			} else {
				log.Debug("reader producer: reader closed")
			}
			close(out)
			close(errs)
			wg.Done()
		}()

		buf := make([]byte, 1500)
		for {
			n, err := r.Read(buf)
			if err != nil {
				errs <- errors.Wrap(err, "could not read")
				break
			}

			select {
			case out <- buf[0:n]:
			case <-ctx.Done():
				log.Debugf("reader producer: %v", ctx.Err())
				return
			}
		}
	}()
	return out, errs
}
