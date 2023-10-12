package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	num := []int{4, 0, 5, -5, 3, 3, 0, -4, -5}
	fmt.Println(threeSumClosest(num, -2))
}
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	res := math.MaxInt32
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, n-1
		for left < right {
			total := nums[i] + nums[left] + nums[right]
			fmt.Println(total, "i:", i, "left:", left, "right:", right)
			fmt.Println(nums[i], "-", nums[left], "*", nums[right])
			if total < target {
				left++
			} else {
				right--
			}
			if math.Abs(float64(target-res)) > math.Abs(float64(target-total)) {
				res = total
			}
			fmt.Println(res)
		}
	}
	return res
}
