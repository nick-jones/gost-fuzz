package main

import (
	"log"
	"net"
)

func main() {
	if ip := net.ParseIP("127.0.0.1"); ip != nil {
		log.Panicln("ok")
	}
}
