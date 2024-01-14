package main

import "fmt"

func main() {
	arr := []int{0, 0, 0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(arr))
}

func removeDuplicates(nums []int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)

		}
		fmt.Println(nums)
	}
	return nums

}
