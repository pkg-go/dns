package main

import (
	"flag"
	"log"
	"net"

	"github.com/go-rfc/dns"
	"github.com/go-rfc/dns/debug"
)

var (
	port = flag.Int("port", 53, "binding port.")
)

func main() {
	flag.Parse()

	addr := &net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: *port}
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		log.Panicf("server: cannot connect: %s", err)
	}
	defer conn.Close()
	log.Printf("server listening on %s", addr)

	data := make([]byte, 576)
	for {
		n, peer, _ := conn.ReadFromUDP(data)
		r := dns.NewReader(data[:n])
		msg := r.ReadMessage()
		debug.PrintMessage(msg)
		conn.WriteToUDP(data, peer)
	}
}
