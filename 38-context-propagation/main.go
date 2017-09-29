package main

import (
	"context"
	"fmt"
	"log"
	"sync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Wait all goroutines are done
	var wg sync.WaitGroup
	wg.Add(4)
	out1 := gen(ctx, &wg)
	out2 := gen(ctx, &wg)
	out3 := gen(ctx, &wg)
	output := merge(ctx, &wg, out1, out2, out3)

	for n := range output {
		fmt.Println(n)
		if n == 10 {
			// propagate cancel to all goroutines
			cancel()
			break
		}
	}
	wg.Wait()
}

func gen(ctx context.Context, wg *sync.WaitGroup) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		defer wg.Done()

		var n int
		for {
			select {
			case ch <- n:
				n++
			case <-ctx.Done():
				log.Printf("gen: %v", ctx.Err())
				return
			}
		}
	}()
	return ch
}

func merge(ctx context.Context, wg *sync.WaitGroup, inputs ...<-chan int) <-chan int {
	var lwg sync.WaitGroup
	outCh := make(chan int)

	// Start an receiver goroutine for each input channel. receiver
	// copies values from ch to outCh until ch is closed, then calls wg.Done.
	receiver := func(ch <-chan int) {
		defer lwg.Done()
		defer func() {
			log.Printf("receiver: %v", ctx.Err())
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
