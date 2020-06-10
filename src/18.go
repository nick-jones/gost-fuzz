package main

import "fmt"

type foo struct {
	str string
}

func main() {
	x := &foo{"banana"}
	fmt.Println(x)
}