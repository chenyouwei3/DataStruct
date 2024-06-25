package main

import "fmt"

func main() {
	arr := []int{3, 4, 5, 1, 2}

	fmt.Println(findMin(arr))
}

func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		pivot := low + (high-low)/2
		if nums[pivot] < nums[high] {
			high = pivot
		} else {
			low = pivot + 1
		}
	}
	return nums[low]
}
