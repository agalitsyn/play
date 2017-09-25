package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net"

	"golang.org/x/net/ipv4"
)

const (
	network         = "eth0"
	mcAddr          = "224.0.0.0:1000"
	src             = "10.50.100.200"
	group           = "224.0.0.0"
	maxDatagramSize = 1500
)

func main() {
	c, err := net.ListenPacket("udp", mcAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	p := ipv4.NewPacketConn(c)
	ifi, err := net.InterfaceByName(network)
	if err != nil {
		log.Fatal(err)
	}

	g := net.UDPAddr{IP: net.ParseIP(group)}
	s := net.UDPAddr{IP: net.ParseIP(src)}
	if err := p.JoinSourceSpecificGroup(ifi, &g, &s); err != nil {
		log.Fatal(err)
	}

	for {
		b := make([]byte, maxDatagramSize)
		n, _, src, err := p.ReadFrom(b)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}

		fmt.Println(n, "bytes read from", src)
		fmt.Println()
		fmt.Println(hex.Dump(b[:n]))
	}
}
