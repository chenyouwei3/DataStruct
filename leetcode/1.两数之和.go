package main

func twoSum(nums []int, target int) []int {
	var bum []int
	for i := range nums {
		for j := range nums {

			sum := nums[i] + nums[j]
			if sum == target && i != j {
				bum = []int{i, j}
				return bum
			}
		}
	}
	return nil
}
