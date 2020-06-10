package main

import "fmt"

func main() {
	list := []string{
		"banana",
		"apple",
		"pears",
	}
	fmt.Printf("the list: %v", list)
	for i := 0; i < 100; i++ {
		fmt.Println(list)
	}
}
