package main

func main() {
	arr := []int{1, 1, 1}
	removeDuplicates2(arr)
}

func removeDuplicates2(nums []int) int {
	numsMap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]]++
		if numsMap[nums[i]] > 2 {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}
