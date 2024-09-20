package main

import (
	"fmt"
	"math"
)

func main() {
	num := "   -42"
	fmt.Println(num)
}

func myAtoiTest0(s string) int {
	result, i, n := 0, 0, len(s)
	sign := 1
	for ; i < n && s[i] == ' '; i++ {
	}
	if i >= n {
		return 0
	}
	switch s[i] {
	case '+':
		i++
	case '-':
		i++
		sign = -1
	}
	for ; i < n; i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
		result = result*10 + int(s[i]-'0')
		if sign*result < math.MinInt32 {
			return math.MinInt32
		}
		if sign*result > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return sign * result
}

func myAtoiTest1(s string) int {
	var total, test1, test2 int
	boll := false
	for _, char := range s {
		if (char < '0' || char > '9') && char != ' ' && char != '-' && char != '+' { //判断是不是那几个东西
			if boll {
				total = -total
			}
			return total
		}
		if test2 > 0 && (char == '-' || char == '+') || test1 > 1 || (char == ' ' && test1 > 0) {
			if boll {
				total = -total
			}
			return total
		}
		if boll {
			total = -total
		}
		if char >= '0' && char <= '9' {
			if total*10+int(char-'0') <= math.MinInt32 || total*10+int(char-'0') >= math.MaxInt32+1 {
				if total > 0 {
					return math.MaxInt32
				}
				return math.MinInt32
			}
			total = total*10 + int(char-'0')
			test2++
		}
		if char == '-' {
			boll = true
			test1++
		}
		if char == '+' {
			test1++
		}
		fmt.Println(boll, total)
	}
	if boll {
		total = -total
	}
	return total
}
