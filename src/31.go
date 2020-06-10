package main

import (
	"log"
	"os"
	"unicode/utf8"
)

func main() {
	if utf8.ValidString(os.Args[1]) {
		log.Println("utf-8!")
	}
}
