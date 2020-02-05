package main

import "fmt"

// 最大子序和
func maxSubArray(nums []int) int {
	maxp := nums[0]
	current := nums[0]
	for i := 1; i < len(nums); i++ {
		current = max(nums[i], current+nums[i])
		maxp = max(maxp, current)
	}
	return maxp
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
