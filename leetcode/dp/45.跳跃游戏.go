package main

import "fmt"

func main() {
	arrary := []int{2, 3, 0, 1, 4}
	fmt.Println(jump(arrary))
}

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	var needChoose, realChoose, step int
	for i, v := range nums {
		if i+v > realChoose {
			realChoose = i + v
			if realChoose >= len(nums)-1 {
				return step + 1
			}
		}
		if i == needChoose {
			needChoose = realChoose
			step++
		}
	}
	return step
}
