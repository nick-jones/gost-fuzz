package main

import (
	"log"
	"os"
	"unicode/utf8"
)

func main() {
	if utf8.ValidString(os.Args[1] + "\t\n\x00\x01\u0002") {
		log.Println("utf-8!")
	}
}
