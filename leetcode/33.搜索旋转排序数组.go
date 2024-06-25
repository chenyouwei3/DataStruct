package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(arr)
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		// 判断哪边是有序的
		if nums[left] <= nums[mid] {
			// 左半边有序
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 右半边有序
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
