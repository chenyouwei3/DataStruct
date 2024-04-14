package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 3, 5, 6}
	fmt.Println(searchInsert1(arr, 7))
}

func searchInsert1(nums []int, target int) int {
	min, max := 0, len(nums)-1
	for min <= max {
		mid := (min + max) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			max--
		} else if nums[mid] < target {
			min++
		}
	}
	var res int
	for i, n := range nums {
		fmt.Println(i, n, target)
		if n >= target {
			res = i
			break
		}
		if i == len(nums)-1 {
			return len(nums)
		}
	}
	return res
}

func searchInsert2(nums []int, target int) int {
	min, max := 0, len(nums)
	for min <= max {
		mid := (min + max) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			max--
		} else if nums[mid] < target {
			min++
		}
	}
	//index := sort.SearchInts(nums, target)
	return sort.SearchInts(nums, target)
}
