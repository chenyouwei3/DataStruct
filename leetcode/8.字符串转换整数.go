package main

import (
	"fmt"
	"math"
)

func main() {
	num := "9223372036854775808"
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

func myAtoiTest(s string) int {
	var number, total, test1, test2 int
	var boll bool
	for _, char := range s {
		//判断是不是那几个东西
		if (char < '0' || char > '9') && char != 32 && char != 45 && char != 43 {
			if boll {
				total = -total
			}

			return total
		}
		if test2 > 0 && (char == 45 || char == 43) || test1 > 1 || (char == 32 && test1 > 0) {
			if boll {
				total = -total
			}

			return total
		}
		if char >= '0' && char <= '9' {
			number = int(char - '0')

			total = total*10 + number
			test2++
			fmt.Println(total, "suanfa")
		}
		if char == 45 {

			boll = true
			test1++
		}
		if char == 43 {
			test1++
		}
		fmt.Println(total)
	}
	fmt.Println(total, "TOAL")
	if boll {
		total = -total
	}
	if total <= -2147483648 || total >= 2147483647 {
		if total > 0 {
			return 2147483647
		}
		return -2147483648
	}
	return total
}
