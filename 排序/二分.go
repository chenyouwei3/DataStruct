package main

import "fmt"

func main() {
	arr := []int{0, 5, 4, 9, 7, 8, 3}
	fmt.Println(BinarySearch(arr, 5))
}

func BinarySearch(a []int, v int) int {
	length := len(a)
	min := 0
	max := length - 1
	for min <= max {
		mid := min + (max-min)/2
		if a[mid] == v {
			return mid
		} else if a[mid] < v {
			min = mid + 1
		} else if a[mid] > v {
			max = mid - 1
		}
	}
	return -1
}
