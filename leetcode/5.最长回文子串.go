package main

func main() {
	string1 := "pwwkew"
	longestPalindrome(string1)
}

func longestPalindrome(s string) string {
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
		for i := 0; i < n-L+1; i++ {
			j := i + L - 1
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
