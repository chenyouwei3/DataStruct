package main

import "fmt"

func main() {
	testDemo := make([][]int, 4)
	testDemo[0] = []int{1, 4}
	testDemo[1] = []int{0, 1}
	testDemo[2] = []int{8, 10}
	testDemo[3] = []int{10, 18}
	fmt.Println(merge(testDemo))
}

func mergeTest(intervals [][]int) [][]int {
	result := make([][]int, len(intervals))
	for i, v := range intervals {
		result[i] = v
	}
	var left, right int
	for i, v := range intervals {
		if (v[0] <= right && right <= v[1]) || (left <= v[0] && v[0] <= right) {
			newSlice := []int{mergeMin(mergeMin(v[0], right), left), mergeMax(mergeMax(v[0], v[1]), right)}
			result = append(result[:i], result[i+1:]...)
			result[i-1] = newSlice
			fmt.Println(result)
		}
		//更新元素
		left, right = v[0], v[1]
		//fmt.Println("left:", left, "right", right, "index", i)
	}
	//fmt.Println("result:", result)
	return result
}

func mergeMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mergeMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func merge(intervals [][]int) [][]int {
	result := make([][]int, len(intervals))
	for i, v := range intervals {
		result[i] = v
	}
	var left, right, i int
	for _, v := range intervals {
		if (v[0] <= right && right <= v[1]) || (left <= v[0] && v[0] <= right) || v[0] == left {
			newSlice := []int{mergeMin(mergeMin(v[0], right), left), mergeMax(mergeMax(v[0], v[1]), right)}
			result = append(result[:i], result[i+1:]...)
			result[i-1] = newSlice
			i--
		}
		i++
		//更新元素
		left, right = v[0], v[1]
		//fmt.Println("left:", left, "right", right, "index", i)
	}
	//fmt.Println("result:", result)
	return result
}
