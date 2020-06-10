package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if strings.EqualFold(os.Args[1], "banana") {
		fmt.Println("fruity")
	}
}
