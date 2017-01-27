package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var file string
	if len(os.Args) > 1 {
		file = os.Args[1]
	} else {
		log.Fatal("file name is required")
	}

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err.Error())
	}

	b, err := byteCounter(f)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("bytes: %v\n", b)

	_, err = f.Seek(0, 0)
	if err != nil {
		log.Fatal(err.Error())
	}

	l, err := lineCounter(f)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("lines: %v\n", l)

	fmt.Println()

	half := int64(b / 2)
	fmt.Printf("half bytes: %v\n", half)
	fmt.Printf("half lines: %v\n", l/2)

	fmt.Println()
	_, err = f.Seek(half, 0)
	if err != nil {
		log.Fatal(err.Error())
	}

	l, err = lineCounter(f)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("lines read after seek half bytes: %v\n", l)

	fmt.Println()
	_, err = f.Seek(0, 2)
	if err != nil {
		log.Fatal(err.Error())
	}

	buf := make([]byte, half)
	c, err := f.ReadAt(buf, half)
	if err != nil && err != io.EOF {
		log.Fatal(err.Error())
	}
	fmt.Printf("bytes read afrer counting byte offset: %v\n", c)

	l = bytes.Count(buf, []byte{'\n'})
	fmt.Printf("lines read after counting byte offset: %v\n", l)
}

func lineCounter(r io.Reader) (int, error) {
	return counter(r, []byte{'\n'})
}

func byteCounter(r io.Reader) (int, error) {
	return counter(r, []byte{})
}

func counter(r io.Reader, sep []byte) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], sep)

		switch {
		case err == io.EOF:
			return count, nil
		case err != nil:
			return count, err
		}
	}
}
