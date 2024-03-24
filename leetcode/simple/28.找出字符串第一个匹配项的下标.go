package main

import "fmt"

func main() {
	haystack := "sadbutsad"
	needle := "sad"
	strStr(haystack, needle)
}

func strStr(haystack string, needle string) int {
	left := 0
	length := len(needle)
	// && left+length <= len(haystack)
	for left < len(haystack) && left+length <= len(haystack) {
		fmt.Println(left+length, haystack[left:left+length])
		if haystack[left:left+length] == needle {

			return left
		}
		left++
	}
	return -1
}
