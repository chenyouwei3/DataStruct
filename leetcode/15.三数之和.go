package main

import (
	"fmt"
	"sort"
)

func main() {
	num := []int{-1, 0, 1, 2, -1, -4}
	threeSum2(num)
}

// 双指针
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := [][]int{}
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			total := nums[i] + nums[left] + nums[right]
			fmt.Println("left:", left, "right:", right, "i:", i)
			fmt.Println(total)
			switch {
			case total < 0:
				left++
			case total > 0:
				right--
			case total == 0:
				res = append(res, []int{nums[i], nums[left], nums[right]})
				fmt.Println("==0", res)
				/*  去重复指针*/
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				/*  去重复指针*/
				left++
				right--
			}
		}
	}
	fmt.Println(res, "last")
	return res
}

func threeSum(nums []int) [][]int {
	results := [][]int{}
	n := len(nums)
	sort.Ints(nums) // 排序使得相同的元素相邻
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // 去除重复的第一个元素
		}
		seen := make(map[int]bool)
		for j := i + 1; j < n; j++ {
			complement := -nums[i] - nums[j]
			if seen[complement] {
				results = append(results, []int{nums[i], complement, nums[j]})
				for j < n-1 && nums[j] == nums[j+1] {
					j++ // 去除重复的第二个元素
				}
			}
			seen[nums[j]] = true
		}
	}

	return results
}
