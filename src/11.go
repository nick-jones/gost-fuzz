package main

import (
	"fmt"
	"os"
)

func main() {
	if check(os.Args[1]) {
		fmt.Println("fruity!")
	}
}

func check(str string) bool {
	return str == "banana"
}