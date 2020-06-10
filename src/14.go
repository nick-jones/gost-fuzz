package main

import (
	"fmt"
)

func main() {
	fmt.Println(double("banana"))
}

func double(str string) string {
	return str + str
}