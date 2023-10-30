package main

import "fmt"

func main() {
	string1 := "chenyouwei3"
	lengthOfLongestSubstring0(string1)
	fmt.Println(lengthOfLongestSubstring0(string1))
}

// 滑动窗口
func lengthOfLongestSubstring0(s string) int {
	if len(s) == 0 {
		return 0
	}
	var freq [127]int
	result, left, right := 0, 0, -1
	for left < len(s) {
		//右指针右移,判断右移后字符是否错误
		if right+1 < len(s) && freq[s[right+1]] == 0 {
			freq[s[right+1]]++
			right++
		} else {
			freq[s[left]]--
			left++
		}
		result = max(result, right-left+1)
	}
	return result
}

func lengthOfLongestSubstringTest(s string) int {
	maxLen, left := 0, 0 // 初始化最长子串长度和窗口起始位置

	charMap := make(map[byte]int) // 用来记录字符最后一次出现的索引

	for right := 0; right < len(s); right++ { // 从头到尾遍历字符串
		if index, ok := charMap[s[right]]; ok && index >= left { // 如果当前字符已经出现过且在当前窗口内
			left = index + 1 // 更新窗口起始位置为相同字符上次出现的位置的下一个位置
		}
		fmt.Printf("%c->", 101)
		charMap[s[right]] = right          // 更新当前字符的索引到 charMap 中
		maxLen = max(maxLen, right-left+1) // 更新最长子串长度
		fmt.Println(charMap)
	}
	fmt.Println(maxLen)

	return maxLen // 返回最长子串长度
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
