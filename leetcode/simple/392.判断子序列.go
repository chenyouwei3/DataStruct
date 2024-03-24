package main

import "fmt"

func main() {
	s := "abc"
	t := "abcddasdsa"
	fmt.Println(isSubsequence(s, t))
}

func isSubsequence(s string, t string) bool {
	for len(t) > 0 && len(s) > 0 {
		if s[0] == t[0] {
			s = s[1:]
		}
		t = t[1:]
	}
	return len(s) == 0
}
