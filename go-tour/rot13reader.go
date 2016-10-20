// https://tour.golang.org/methods/23
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(c byte) byte {
	// c = ASCII decimal code
	// See algorithm in https://en.wikipedia.org/wiki/ROT13
	if c >= 'a' && c <= 'm' || c >= 'A' && c <= 'M' {
		c += 13
	} else if c >= 'n' && c <= 'z' || c >= 'N' && c <= 'Z' {
		c -= 13
	}

	return c
}

func (self rot13Reader) Read(bstream []byte) (readed int, err error) {
	readed, err = self.r.Read(bstream)
	for i := 0; i < readed; i++ {
		bstream[i] = rot13(bstream[i])
	}

	return readed, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}