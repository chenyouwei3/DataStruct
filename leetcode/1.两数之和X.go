package main

import "fmt"

func main() {
	num := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(num, target))
}

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

func twoSumTest0(nums []int, target int) []int {
	charMap := map[int]int{}
	for k, v := range nums {
		if p, ok := charMap[target-v]; ok {
			fmt.Println(p, k)
			return []int{p, k}
		}
		charMap[v] = k
	}
	return nil
}

func twoSumTest1(nums []int, target int) []int {
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
