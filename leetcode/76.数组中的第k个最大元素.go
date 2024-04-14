package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(arr, 2))
}

func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	index := len(nums) - k
	return nums[index]
}
