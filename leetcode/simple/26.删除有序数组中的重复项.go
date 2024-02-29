package main

import "fmt"

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicatess(arr))
}

func removeDuplicates(nums []int) int {
	last, res := 0, 0
	for last < len(nums)-1 {
		for nums[last] == nums[res] {
			res++
		}
		nums[last+1] = nums[res]
		last++
	}
	return last + 1
}

func removeDuplicatess(nums []int) int {
	mmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mmap[nums[i]]++
		if mmap[nums[i]] >= 2 {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	fmt.Println(nums)
	return len(nums)
}
