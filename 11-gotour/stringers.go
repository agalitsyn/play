// https://tour.golang.org/methods/18
package main

import (
	"bytes"
	"fmt"
)

type IPAddr [4]byte

func (self IPAddr) String() string {
	var buffer bytes.Buffer

	// Of course we can do `return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])`, but it's not true way
	for i := 0; i < len(self); i++ {
		buffer.WriteString(fmt.Sprintf("%v.", self[i]))
	}

	return buffer.String()
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for key, value := range addrs {
		fmt.Printf("%v: %v\n", key, value)
	}
}