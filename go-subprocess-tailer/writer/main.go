package main

import (
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var logFile string
	if len(os.Args) > 0 {
		logFile = os.Args[1]
	} else {
		log.Fatal("log file name is required")
	}

	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer f.Close()

	for {
		data := RandStringBytes(64)
		log.Println(string(data))

		n, err := f.Write(data)
		if err == nil && n < len(data) {
			err = io.ErrShortWrite
		}
		if err != nil {
			log.Fatalf(err.Error())
		}

		_, err = f.WriteString("\n")
		if err != nil {
			log.Fatalf(err.Error())
		}

		f.Sync()
	}
}

func RandStringBytes(n int) []byte {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}
