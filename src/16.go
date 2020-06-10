package main

import "fmt"

type foo struct {
	str string
}

func main() {
	fmt.Println(&foo{
		str: "banana",
	})
}