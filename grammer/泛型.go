package main

import "fmt"

func main() {
	Min(3.00, 5.0)
	Max(5, 3)
	var a INTSlice[int]
	a := []int{1, 2, 6, 4, 5}
}

type INTSlice[T int | string] []T

func Min[T int | float64](a, b T) {
	if a > b {
		fmt.Println(b)
	}
	go fmt.Println(a)
}

func Max[T int | float64](a, b T) {
	if b > a {
		fmt.Println(b)
	}
	fmt.Println(a)
}
