package main

import "fmt"

func main() {
	str := "chenyouwei3"
	fmt.Println(test(str))
}

func test(s string) int {
	left, right, res := 0, 0, 0
	buckets := make(map[byte]int, len(s))
	for left < len(s) {
		if index, ok := buckets[s[left]]; ok && index >= right {
			right = index + 1
		}
		buckets[s[left]] = left
		left++
	}

}
