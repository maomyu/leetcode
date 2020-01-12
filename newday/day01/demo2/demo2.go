package main

import "fmt"

// 买卖股票的最佳时机 II

func maxProfit(prices []int) int {
	max := 0

	for j := 0; j < len(prices)-1; j++ {
		if prices[j] < prices[j+1] {
			max = max + (prices[j+1] - prices[j])
		}
	}

	return max

}
func main() {
	data := []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(data))
}
