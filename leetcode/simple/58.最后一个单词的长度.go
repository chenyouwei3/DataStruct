package main

import "fmt"

func main() {
	s := "   fly me   to   the moon  "
	fmt.Println("res", lengthOfLastWord1(s))
}

func lengthOfLastWord(s string) int {
	lengthInedex := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == ` ` {
			lengthInedex = i
		}
	}
	fmt.Println(s[lengthInedex+1:])
	return len(s[lengthInedex+1:])
}

func lengthOfLastWord1(s string) int {
	last := len(s) - 1
	for last >= 0 && s[last] == ' ' {
		last--
	}
	if last < 0 {
		return 0
	}
	fmt.Println("last:", last, string(s[last]))
	first := last
	for first >= 0 && s[first] != ' ' {
		first--
	}
	fmt.Println("first:", first, string(s[first+1]))
	return last - first
}
