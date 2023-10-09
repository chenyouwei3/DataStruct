package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(num[2:])
	num = append(num[:1], num[2:]...)
	fmt.Println(num)
}
