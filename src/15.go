package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(double(os.Args[1]))
}

func double(str string) string {
	return str + str
}