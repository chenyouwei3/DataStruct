package main

import (
	"fmt"
)

func main() {
	x := -15344564
	fmt.Println(reverse(x))
}

func reverse(x int) int {
	var sum int
	for x != 0 {
		if temp := int32(sum); (temp*10)/10 != temp {
			return 0
		}
		sum = sum*10 + x%10
		x = x / 10
	}
	return sum
}

func reverseTest(x int) int {
	if x%10 == 0 {
		return isInt64(x / 10)
	}
	return isInt64(x)
}

func isInt64(x int) int {
	var sum int
	for x != 0 {
		sum = sum*10 + x%10
		x /= 10
	}
	if sum < -2147483648 || sum > 2147483647 {
		return 0
	}
	return sum
}
