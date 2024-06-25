package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive1(arr))
}
func longestConsecutive1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	sort.Ints(nums)
	index, res := 0, 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			res = maxTest0(res, i-index)
			continue
		} else if nums[i] == nums[i-1] {
			index++
		} else {
			index = i
		}
	}
	return res + 1
}

func maxTest0(a, b int) int {
	if a > b {
		return a
	}
	return b
}
