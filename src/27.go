package main

import (
	"log"
	"net"
)

func main() {
	if ip := net.ParseIP("banana"); ip != nil {
		log.Panicln("ok")
	}
}
