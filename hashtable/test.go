package main

import "fmt"

func main() {
	mmap := make(map[int]int)
	mmap[0] = 0
	mmap[1] = 1
	mmap[2] = 2
	mmap[3] = 3
	fmt.Println(mmap)
}
