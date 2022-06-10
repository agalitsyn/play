package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/kamilsk/retry/v5"
	"github.com/kamilsk/retry/v5/strategy"
)

func main() {
	fmt.Println(waitForPodDeletion(context.Background(), 1*time.Second, 10*time.Second))
}

func waitForPodDeletion(ctx context.Context, delay, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	action := func(ctx context.Context) (err error) {
		log.Println("poll")
		err = deletePod()
		if err != nil && err == ErrNotFound {
			return nil
		}
		return err
	}

	how := retry.How{
		strategy.Limit(5),
		strategy.Wait(delay),
	}

	log.Println("start")
	return retry.Do(ctx, action, how...)
}

var ErrNotFound = errors.New("not found")

func deletePod() error {
	time.Sleep(250 * time.Millisecond)
	return ErrNotFound
}
