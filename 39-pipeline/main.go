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

	signalChan := make(chan os.Signal, 1)
	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup

	// wg.Add(4)
	// rawData, readerErrs := readerProducer(ctx, &wg, bytes.NewReader([]byte("foo bar baz")))
	// rawData2, readerErrs2 := readerProducer(ctx, &wg, badReader{})
	wg.Add(2)
	rawData3 := randomProducer(ctx, &wg)
	rawData4 := randomProducer(ctx, &wg)

	// wg.Add(1)
	// errs := mergeErrors(ctx, &wg, readerErrs, readerErrs2)

	wg.Add(2)
	// processedData := upperCaseProcessor(ctx, &wg, rawData)
	// processedData2 := upperCaseProcessor(ctx, &wg, rawData2)
	processedData3 := upperCaseProcessor(ctx, &wg, rawData3)
	processedData4 := upperCaseProcessor(ctx, &wg, rawData4)

	wg.Add(1)
	output := merge(ctx, &wg, processedData3, processedData4)

	// go func() {
	// 	s := <-signalChan
	// 	log.Printf("captured %v, exiting", s)
	// 	cancel()
	// }()
	// go func() {
	// 	for err := range errs {
	// 		log.Error(err)
	// 	}
	// 	log.Info("errors done")
	// }()

	for data := range output {
		fmt.Println(string(data))
		cancel()
	}

	wg.Wait()
}

func upperCaseProcessor(ctx context.Context, wg *sync.WaitGroup, in <-chan []byte) <-chan []byte {
	out := make(chan []byte)
	go func() {
		defer wg.Done()
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

func merge(ctx context.Context, wg *sync.WaitGroup, inputs ...<-chan []byte) <-chan []byte {
	var lwg sync.WaitGroup
	out := make(chan []byte)

	output := func(c <-chan []byte) {
		defer lwg.Done()

		for d := range c {
			select {
			case out <- d:
			case <-ctx.Done():
				log.Debug("cancel merge receiver")
				return
			}
		}
	}

	lwg.Add(len(inputs))
	for _, input := range inputs {
		go output(input)
	}

	go func() {
		defer wg.Done()

		lwg.Wait()
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

func randomProducer(ctx context.Context, wg *sync.WaitGroup) <-chan []byte {
	out := make(chan []byte)
	go func() {
		defer close(out)
		defer wg.Done()

		for {
			select {
			case out <- RandBytesString(8):
			case <-ctx.Done():
				log.Debug("cancel random producer")
				return
			}
		}
	}()
	return out
}

func readerProducer(ctx context.Context, wg *sync.WaitGroup, r io.Reader) (chan []byte, chan error) {
	out := make(chan []byte)
	errs := make(chan error)
	go func() {
		defer close(out)
		defer close(errs)
		defer wg.Done()

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

func mergeErrors(ctx context.Context, wg *sync.WaitGroup, inputs ...<-chan error) <-chan error {
	var lwg sync.WaitGroup
	outCh := make(chan error)

	receiver := func(ch <-chan error) {
		defer lwg.Done()

		for e := range ch {
			select {
			case outCh <- e:
			case <-ctx.Done():
				outCh <- errors.Wrap(ctx.Err(), "cancel merging errors")
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
