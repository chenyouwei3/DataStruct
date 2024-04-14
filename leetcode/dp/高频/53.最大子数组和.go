package main

import "fmt"

func main() {
	arrary := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(arrary))
}

func maxSubArray(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	dp, res := make([]int, length), nums[0]
	dp[0] = res
	for i := 1; i < length; i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		res = maxSubArrayAuxiliary(dp[i], res)
	}
	return res
}

func maxSubArrayAuxiliary(a, b int) int {
	if a > b {
		return a
	}
	return b
}
