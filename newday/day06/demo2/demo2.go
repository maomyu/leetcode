package main

import "fmt"

// 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

// 如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。

// 注意你不能在买入股票前卖出股票。

// 暴力法
func maxProfit(prices []int) int {
	i, j := 0, len(prices)-1
	max := 0
	for i < len(prices) {
		for j > i {
			temp := prices[j] - prices[i]
			if max < temp {
				max = temp
			}
			j--
		}
		j = len(prices) - 1
		i++
	}
	return max
}

// 谷峰法
func maxProfit_two(prices []int) int {
	minprices := 1 << 31
	maxprofites := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minprices {
			minprices = prices[i]
		} else if prices[i]-minprices > maxprofites {
			maxprofites = prices[i] - minprices
		}
	}
	return maxprofites
}
func main() {
	prices := []int{10, 2, 9, 1, 2, 1, 3, 1}
	fmt.Println(maxProfit_two(prices))
}
