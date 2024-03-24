package main

import "fmt"

func main() {
	nums := []int{-2}
	fmt.Println(maxProduct(nums))
}

func maxProduct(nums []int) int {
	maxTemp, minTemp, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxTemp, minTemp = minTemp, maxTemp
		}
		maxTemp = auxiliaryMaxProduct(nums[i], maxTemp*nums[i])
		minTemp = auxiliaryMinProduct(nums[i], minTemp*nums[i])
		res = auxiliaryMaxProduct(maxTemp, res)
	}
	return res
}

func auxiliaryMaxProduct(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func auxiliaryMinProduct(a, b int) int {
	if a < b {
		return a
	}
	return b
}
