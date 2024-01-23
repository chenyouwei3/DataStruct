package main

import "fmt"

func main() {
	isPalindromeNum(121)
	fmt.Println(isPalindromeNum(121))
}

func isPalindromeNum(x int) bool {
	y := x
	if x%10 == 0 || x < 0 {
		return false
	}
	var sum int
	for x != 0 {
		sum = sum*10 + x%10
		x /= 10
		fmt.Println(sum)
	}
	if sum == y {
		return true
	}
	return false
}
