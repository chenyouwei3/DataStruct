package main

import (
	"fmt"
)

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}

func letterCombinations(digits string) (res []string) {
	charString := [...]string{" ", " ", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	n := len(digits)
	str := make([]byte, n)
	if n == 0 {
		return
	}
	var strFuncion func(int)
	strFuncion = func(i int) {
		if i == n {
			res = append(res, string(str))
			return
		}
		for _, v := range charString[digits[i]-48] {
			str[i] = byte(v)
			strFuncion(i + 1)
		}
	}
	strFuncion(0)
	return
}
