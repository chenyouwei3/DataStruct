package main

import "fmt"

func main() {
	str := "babad"
	fmt.Println(str[0])
	fmt.Println(longestPalindromeDemo(str))
}

/*动态规划*/
func longestPalindromeDemo(s string) string {
	dp := make([][]bool, len(s))
	result := s[0:1] //初始化结果(最小的回文就是单个字符)
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true // 根据case     1 初始数据
	}
	for length := 2; length <= len(s); length++ { //长度固定，不断移动起点
		for start := 0; start <= len(s)-length; start++ { //长度固定，不断移动起点
			end := start + length - 1
			if s[start] != s[end] { //首尾不同则不可能为回文
				continue
			} else if length < 3 {
				dp[start][end] = true //即case 2的判断
			} else {
				dp[start][end] = dp[start+1][end-1] //状态转移
			}
			if dp[start][end] && (end-start) >= len(result) { //记录最大值
				result = s[start : end+1]
			}
		}
	}
	return result
}

// 中心扩散，时间复杂度 O(n^2)，空间复杂度 O(1)
func longestPalindrome2(s string) string {
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

/*判断回文子串*/
func isPalindromeString(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

/*求字符串的子串*/
func stringOfSon(s string) []string {
	var result []string
	stringOfSon2(s, "", 0, &result)
	return result
}

func stringOfSon2(s, currentSubset string, index int, result *[]string) {
	if index == len(s) {
		*result = append(*result, currentSubset)
		return
	}
	// 不包含当前字符的情况
	stringOfSon2(s, currentSubset, index+1, result)
	// 包含当前字符的情况
	stringOfSon2(s, currentSubset+string(s[index]), index+1, result)
}
