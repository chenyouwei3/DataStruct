package main

func main() {

}

// maxSubArrayBruteForce 使用暴力法计算连续子数组的最大和
func maxSubArrayBruteForce(nums []int) int {
	if len(nums) == 0 {
		// 根据需求返回 0 或其他值
		return 0
	}
	maxSum := nums[0]
	for i := 0; i < len(nums); i++ {
		currentSum := 0
		for j := i; j < len(nums); j++ {
			currentSum += nums[j]
			if currentSum > maxSum {
				maxSum = currentSum
			}
		}
	}
	return maxSum
}

func maxSubArrayKadane(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	currentSum := nums[0]
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum+nums[i])
		maxSum = max(maxSum, currentSum)
	}

	return maxSum
}

// max 是一个辅助函数，返回两个整数中的较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
