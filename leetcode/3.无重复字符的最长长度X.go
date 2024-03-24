package main

import "fmt"

func main() {
	//string1 := "pwwkew"
	string2 := "chenyouwei3"
	//lengthOfLongestSubstring0(string2)
	fmt.Println(lengthOfLongestSubstring0(string2))
}

/*滑动窗口-哈希桶*/
func lengthOfLongestSubstring0(s string) int {
	right, left, res := 0, 0, 0           // 初始化右边界、左边界和结果长度为0
	buckets := make(map[byte]int, len(s)) // 创建字典用于记录字符的索引位置
	for left < len(s) {                   // 循环直到左边界到达字符串的末尾
		if index, ok := buckets[s[left]]; ok && index >= right {
			// 如果当前字符在字典中存在，并且索引大于等于右边界
			right = index + 1 // 更新右边界，将其移动到重复字符的下一个位置，形成新的滑动窗口
			fmt.Println("left:", left, "right:", right, "如果存在的化")
		}
		fmt.Println("left:", left, "right:", right)
		buckets[s[left]] = left                            // 将当前字符及其索引位置添加到字典中
		left++                                             // 左边界向右移动
		res = lengthOfLongestSubstringMax(res, left-right) // 更新最长子串的长度
		fmt.Println("res:", res, "left-right", left-right)
	}
	return res // 返回最长不重复子串的长度
}

func lengthOfLongestSubstringMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
