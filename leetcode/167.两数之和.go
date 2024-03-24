package main

import "fmt"

func main() {
	var slice = []int{2, 7, 11, 15}
	fmt.Println(twoSum(slice, 9))
}

func twoSum(numbers []int, target int) []int {
	buckets := make(map[int]int)
	for i, v := range numbers {
		anthor := target - v
		if _, ok := buckets[anthor]; ok {
			return []int{buckets[anthor], i + 1}
		}
		buckets[v] = i + 1
	}
	return nil
}
