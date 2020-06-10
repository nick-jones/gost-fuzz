package main

import "fmt"

type foo struct {
	str1 string
	str2 string
}

func main() {
	x := &foo{
		str1: "banana",
		str2: "apple",
	}
	fmt.Println(x)
}
