package main

import "fmt"

func main() {
	list := []int{1, 3, 2, 5, 4}
	list1 := []int{6, 8, 7, 10, 9}
	list = append(list[:3], list1[3:]...)
	fmt.Println(list)
}
