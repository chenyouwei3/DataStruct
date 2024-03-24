package main

import (
	"strconv"
)

func main() {
	num1 := "-132"
	num2 := "2"
	multiply(num1, num2)
}

func multiply(num1 string, num2 string) string {
	number := multiplyTest(num1) * multiplyTest(num2)
	return strconv.FormatInt(number, 10)
}

func multiplyTest(num string) (total int64) {
	var f int64 = 1
	for i, v := range num {
		if v == '-' {
			f = -1
			continue
		}
		total += multiplyPow(len(num)-1-i) * (int64(v) - 48)
	}
	return total * f
}

func multiplyPow(n int) int64 {
	var m int64 = 1
	for i := 0; i < n; i++ {
		m *= 10
	}
	return m
}
