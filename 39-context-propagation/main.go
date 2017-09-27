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

	out1 := gen(ctx)
	out2 := gen(ctx)
	out3 := gen(ctx)
	output := merge(ctx, out1, out2, out3)

	for n := range output {
		fmt.Println(n)
		if n == 10 {
			log.Println("caceling")
			cancel()
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
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

func merge(ctx context.Context, inputs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	outCh := make(chan int)

	receiver := func(ch <-chan int) {
		defer wg.Done()
		for i := range ch {
			select {
			case outCh <- i:
			case <-ctx.Done():
				log.Printf("close merge: %v", ctx.Err())
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
