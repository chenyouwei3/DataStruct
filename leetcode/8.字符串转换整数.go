package main

import "fmt"

func main() {
	num := "+-12"

	fmt.Println(myAtoi(num), "main")
}

func myAtoi(s string) int {
	var number, total, test int
	var boll bool
	for _, char := range s {
		if (char < '0' || char > '9') && char != 32 && char != 45 && char != 43 {
			return total
		}
		if char >= '0' && char <= '9' {
			number = int(char - '0')
			total = total*10 + number
		}
		if char == 45 {
			boll = true
			test++
		}
		if char == 43 {
			test++
		}
		if test > 1 {
			return 0
		}
	}
	if boll {
		total = -total
	}
	fmt.Println(total)
	if total <= -2147483648 || total >= 2147483647 {
		if total > 0 {
			return 2147483647
		}
		return -2147483648
	}
	return total
}
