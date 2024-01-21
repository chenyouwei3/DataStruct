package main

import "fmt"

func main() {
	resString := "babad"
	fmt.Println(longestPalindrome1(resString))
}

/*滑动窗口*/
func longestPalindrome1(s string) string {
	if len(s) == 0 {
		return ""
	}
	left, right, pl, pr := 0, -1, 0, 0
	for left < len(s) {
		// 移动到相同字母的最右边（如果有相同字母）
		for right+1 < len(s) && s[left] == s[right+1] {
			right++
		}
		// 找到回文的边界
		for left-1 >= 0 && right+1 < len(s) && s[left-1] == s[right+1] {
			left--
			right++
		}
		if right-left > pr-pl {
			pl, pr = left, right
		}
		// 重置到下一次寻找回文的中心
		left = (left+right)/2 + 1
		right = left
	}
	return s[pl : pr+1]
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
