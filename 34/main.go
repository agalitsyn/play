package main

import (
	"errors"
	"log"
)

var shouldThrowErr bool

func main() {
	doneCh := make(chan struct{}, 1)
	errCh := make(chan error, 1)

	go worker(doneCh, errCh)
	shouldThrowErr = true
	go worker(doneCh, errCh)

	select {
	case err := <-errCh:
		log.Println(err)
	case <-doneCh:
	}
}

func worker(doneCh chan struct{}, errCh chan error) {
	log.Println("got task")
	if shouldThrowErr {
		errCh <- errors.New("err!")
	}
	doneCh <- struct{}{}
}
