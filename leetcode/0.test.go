package main

import "fmt"

func main() {
	arr := []int{7, 1, 5, 3, 6, 4}
	maxProfit(arr)
}

func maxProfit(prices []int) int {
	total, last, temp := 0, prices[0], true
	for i := 1; i < len(prices); i++ {
		if !temp {
			//买票
			last = prices[i]
			temp = true
		}
		fmt.Println(last, prices[i], total, temp)
		if prices[i]-last > 0 && temp && i%2 != 0 {
			total = prices[i] - last
			temp = false //没有票了手上
		}

		fmt.Println(total, temp)
	}
	fmt.Println(total)
	return 0
}
