package main

import "fmt"

func main() {
	arrary := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(arrary))
	test(arrary)

}

func test(nums []int) {
	n := len(nums)
	before := make([]int, n)
	before[0] = 1
	for i := 1; i < n; i++ {
		before[i] = before[i-1] * nums[i-1]
	}
	back := 1
	for i := n - 1; i >= 0; i-- {
		before[i] *= back
		back *= nums[i]
	}
	fmt.Println(before)
}

func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	ans[0] = 1
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	suf := 1
	for i := n - 1; i >= 0; i-- {
		ans[i] *= suf
		suf *= nums[i]
	}
	return ans
}
