package main

import (
	"fmt"
)

// 二分查找法——搜索位置
func searchInsert(nums []int, target int) int {
	if len(nums) <= 0 {
		return 0
	}
	low := 0
	high := len(nums) - 1
	mid := 0
	for low <= high {
		mid = (low + high) / 2
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			return mid
		}
	}
	if target > nums[mid] {
		return mid + 1
	}
	return mid
}

func main() {
	nums := []int{1}
	fmt.Println(searchInsert(nums, 2))
}
