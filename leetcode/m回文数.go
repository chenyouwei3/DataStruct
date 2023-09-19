package main

import "fmt"

func main() {
	isPalindrome(121)
	fmt.Println(isPalindrome(121))
}

func isPalindrome(x int) bool {
	y := x
	if x%10 == 0 || x < 0 {
		return false
	}
	var sum int
	for x != 0 {
		sum = sum*10 + x%10
		x /= 10
	}
	if sum == y {
		return true
	}
	return false
}
