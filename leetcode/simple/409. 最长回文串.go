package main

import "fmt"

func main() {
	str := "abccccdd"
	fmt.Println(longestPalindrome(str))
}

func longestPalindrome(s string) int {
	buckets := make(map[int32]int)
	for _, str := range s {
		buckets[str]++
	}
	var i int
	for _, v := range buckets {
		i += v / 2 * 2
		if i%2 == 0 && v%2 == 1 {
			i++
		}

	}
	return i
}
