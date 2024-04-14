package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	index := 2
	element := 10
	insert := []int{element}
	slice1 = append(slice1[:index], append(insert, slice1[index:]...)...)
	fmt.Println(slice1)

}
