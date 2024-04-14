package main

import "fmt"

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(subarraySum(arr, 3))
}

func subarraySum(nums []int, k int) int {
	length := len(nums)
	sun, res := 0, 0
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == k {
				fmt.Println("index:", i, j)
				fmt.Println(nums[i], nums[j])
				res++
			}
		}
	}
	return res
}
