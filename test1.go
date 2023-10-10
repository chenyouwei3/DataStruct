package main

import (
	"fmt"
	"math"
)

func main() {
	num := "  +012 123"

	fmt.Println(myAtoi(num), "main")
}

func myAtoi(s string) int {
	// 去除前导空格
	i := 0
	for i < len(s) && s[i] == ' ' {
		i++
	}
	// 判断正负号
	sign := 1
	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	// 扫描数字字符
	num := 0
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		digit := int(s[i] - '0')
		// 检查溢出
		if num > math.MaxInt32/10 || (num == math.MaxInt32/10 && digit > 7) {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		num = num*10 + digit
		i++
	}
	return num * sign
}
