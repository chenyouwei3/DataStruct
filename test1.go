package main

import "fmt"

func main() {
	s1 := "babad"
	fmt.Println(s1[0:1])
	longestPalindrome1(s1)
}

func longestPalindrome1(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	start, maxLen := 0, 1
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	for L := 2; L <= n; L++ {
		for i := 0; i <= n-L; i++ {
			j := i + L - 1 //
			if s[i] == s[j] && (L == 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				if L > maxLen {
					start = i
					maxLen = L
				}
			}
		}
	}
	return s[start : start+maxLen]
}

func stringTrue(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}
	return false
}
