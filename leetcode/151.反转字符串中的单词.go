package main

import (
	"fmt"
)

func main() {
	s := " "
	fmt.Println("res", reverseWords(s))
}

func reverseWords(s string) string {
	slice := make([][]byte, len(s))
	reverse := make([]string, len(slice))
	var j int
	var temp uint8
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && temp == ' ' {
			j++
		}
		temp = s[i]
		if s[i] == ' ' {
			continue
		}
		if s[i] != ' ' {
			slice[j] = append(slice[j], s[i])
		}
	}
	for i, v := range slice {
		reverse[len(slice)-1-i] = string(v)
	}
	var res string
	for _, s := range reverse {
		if s == "" {
			continue
		}
		res += " " + s
	}
	return res[1:]
}
