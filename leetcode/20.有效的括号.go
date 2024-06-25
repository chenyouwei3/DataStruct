package main

import (
	"fmt"
)

func main() {
	strs := "()[]{}"
	fmt.Println("res", isValid(strs))
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1] //删除一个
			fmt.Println(stack)
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
