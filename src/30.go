package main

import (
	"log"
	"unicode/utf8"
)

const b = `banana`

func main() {
	if utf8.ValidString(b) {
		log.Println("of course")
	}
}
