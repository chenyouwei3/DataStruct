package main

import "fmt"

func main() {
	var testslice = []int{}
	testslice = append(testslice, 1)
	fmt.Println(cap(testslice))
}
