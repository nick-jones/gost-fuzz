package main

import (
	"fmt"
	"os"
)

func main() {
	if os.Args[1] == "banana" {
		fmt.Println("fruity!")
	}
}
