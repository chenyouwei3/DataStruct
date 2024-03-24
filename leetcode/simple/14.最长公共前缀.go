package main

import "fmt"

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix1(strs))
}

func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return " "
	}
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		fmt.Println(string(c))
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			fmt.Println(len(strs[i]), "AND", strs[j][i])
			if i == len(strs[j]) || strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
