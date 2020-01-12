package main

import "fmt"

// 旋转数组
func rotate_first(nums []int, k int) {
	lens := len(nums)
	rnums := nums
	// 带移动的数组
	nums = nums[lens-k:]
	k = lens - k
	i := 0
	for k != 0 {
		nums = append(nums, rnums[i])
		i++
		k--
	}
	fmt.Println(nums)
}
func rotate(nums []int, k int) {
	if k == 0 {
		return
	}
	temp := nums[len(nums)-1]
	for i := len(nums) - 1; i >= 1; i-- {
		nums[i] = nums[i-1]
	}
	nums[0] = temp
	rotate(nums, k-1)

}
func main() {
	data := []int{-1, -100, 3, 99}
	rotate(data, 2)
	fmt.Print(data)
}
