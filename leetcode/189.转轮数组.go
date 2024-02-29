package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	rotate1(arr, 3)
}

func rotate(nums []int, k int) {
	var temp int
	var tempArray []int
	res := &nums
	for i := 0; i < k; i++ {
		temp = nums[len(nums)-1]
		nums = append(nums[:len(nums)-1], nums[len(nums):]...)
		tempArray = []int{temp}
		tempArray = append(tempArray, nums...)
		nums = tempArray
	}
	fmt.Println(res)
}

func rotate1(nums []int, k int) {
	//res := &nums
	n := len(nums)
	k %= n // 防止 k 大于数组长度
	rotateReverse(nums)
	fmt.Println("nums", nums)
	fmt.Println("nums[:k]反转前:", nums[:k])
	rotateReverse(nums[:k])
	fmt.Println("nums[:k]反转后:", nums[:k])
	fmt.Println("nums:", nums)
	fmt.Println("nums[k:]反转前:", nums[k:])
	rotateReverse(nums[k:])
	fmt.Println("nums[k:]反转后:", nums[k:])
	fmt.Println(nums[k:], nums)
}

// 反转数组
func rotateReverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
