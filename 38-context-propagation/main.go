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

	var wg sync.WaitGroup
	wg.Add(4)
	out1 := gen(ctx, &wg)
	out2 := gen(ctx, &wg)
	out3 := gen(ctx, &wg)
	output := merge(ctx, &wg, out1, out2, out3)

	for n := range output {
		fmt.Println(n)
		if n == 10 {
			log.Println("canceling")
			cancel()
			break
		}
	}
	wg.Wait()
}

func gen(ctx context.Context, wg *sync.WaitGroup) <-chan int {
	ch := make(chan int)
	go func() {
		defer wg.Done()

		var n int
		for {
			select {
			case <-ctx.Done():
				log.Printf("close gen: %v", ctx.Err())
				return
			case ch <- n:
				n++
			}
		}
	}()
	return ch
}

func merge(ctx context.Context, wg *sync.WaitGroup, inputs ...<-chan int) <-chan int {
	var lwg sync.WaitGroup
	outCh := make(chan int)

	receiver := func(ch <-chan int) {
		defer lwg.Done()
		for i := range ch {
			select {
			case outCh <- i:
			case <-ctx.Done():
				log.Printf("close receiver: %v", ctx.Err())
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
