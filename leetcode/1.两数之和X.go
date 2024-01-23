package main

import "fmt"

func main() {
	num := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(num, target))
}

/*暴力法*/
func twoSumTest(nums []int, target int) []int {
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

/*哈希桶*/
func twoSum(arrary []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(arrary); i++ {
		another := target - arrary[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[arrary[i]] = i
	}
	return nil
}
