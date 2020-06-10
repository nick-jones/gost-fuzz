package main

import "fmt"

type foo struct {
	str string
}

func main() {
	x := &foo{}
	x.str = "banana"
	fmt.Println(x)
}