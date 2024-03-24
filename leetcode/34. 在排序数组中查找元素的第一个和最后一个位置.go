package main

import "fmt"

func main() {
	num := []int{5}
	fmt.Println(searchRange(num, 5))
}
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] != target {
			left++
			continue
		}
		if nums[right] != target {
			right--
			continue
		}
		return []int{left, right}
	}
	return []int{-1, -1}
}
