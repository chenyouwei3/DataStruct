package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "race a car"
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var tempSlice []byte
	for _, v := range s {
		if ('a' <= v && v <= 'z') || ('0' <= v && v <= '9') {
			tempSlice = append(tempSlice, byte(v))
		}
	}
	left, right := 0, len(tempSlice)-1
	for left <= right {
		if tempSlice[left] != tempSlice[right] {
			return false
		}
		left++
		right--
	}
	return true
}
