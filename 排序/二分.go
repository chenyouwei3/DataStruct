package main

import "fmt"

func main() {
	arr := []int{0, 5, 4, 9, 7, 8, 3}
	fmt.Println(BinarySearch(arr, 9))
}

func BinarySearch(a []int, v int) int {
	min, max := 0, len(a)-1
	for min <= max {
		mid := (max + min) / 2
		if a[mid] == v {
			return mid
		} else if a[mid] > v {
			max--
		} else if a[mid] < v {
			min--
		}
	}
	return -1
}
