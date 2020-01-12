package main

import "fmt"

// 从排序数组中删除重复项
func removeDuplicates(nums []int) int {
	// 【0,0,1,1,1,2,2,3,3,4】
	// 【0,1,0,1,1,2,2,3,3,4】
	// 首指针
	length := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[length] {
			length++
			nums[length] = nums[i]
		}
	}

	return length + 1
}

func main() {
	data := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(data))
}
