// https://tour.golang.org/methods/22
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (self MyReader) Read(bstream []byte) (int, error) {
	size := len(bstream)
	for i := 0; i < size; i++ {
		bstream[i] = 'A'
	}

	return size, nil
}

func main() {
	reader.Validate(MyReader{})
}