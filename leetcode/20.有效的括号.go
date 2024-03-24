package main

import "fmt"

func main() {
	strs := "{[]}"
	fmt.Println("res", isValid(strs))
}

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	fmt.Println(s[len(s)-1])
	bucket := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	buckets := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	var temp string
	for i, v := range s {
		str := s[len(s)-i-1]
		fmt.Println("str", string(str), "v", buckets[string(v)])
		if i%2 == 1 && string(str) != buckets[string(v)] {
			if bucket[string(v)] != temp {
				return false
			}
		}
		temp = string(v)
	}
	return true
}
