package main

import "fmt"

func main() {
	string1 := "chenyouwei3"
	lengthOfLongestSubstring0(string1)
	fmt.Println(lengthOfLongestSubstring0(string1))
}

/*滑动窗口-哈希桶*/
func lengthOfLongestSubstring0(s string) int {
	right, left, res := 0, 0, 0
	buckets := make(map[byte]int, len(s))
	for left < len(s) {
		if index, ok := buckets[s[left]]; ok && index >= right {
			right = index + 1 //滑动窗口,窗口长度为1
			fmt.Println(index, right)
		}
		buckets[s[left]] = left
		left++
		res = lengthOfLongestSubstringMax(res, left-right)
	}
	return res
}

func lengthOfLongestSubstringMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
