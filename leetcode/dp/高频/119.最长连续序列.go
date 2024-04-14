package main

import (
	"fmt"
	"sort"
)

func main() {
	arrary := []int{0, 1, 1, 2}
	fmt.Println(longestConsecutive(arrary))
}

func longestConsecutive(nums []int) int {
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
			res = maxTest(res, i-index)
			continue
		} else if nums[i] == nums[i-1] {
			index++
		} else {
			index = i
		}
	}
	return res + 1
}

func maxTest(a, b int) int {
	if a > b {
		return a
	}
	return b
}
