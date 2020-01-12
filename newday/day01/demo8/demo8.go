package main

import "fmt"

// 移动0到末尾，同时保持非0元素的顺序[0,0,1]
func moveZeroes(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] == 0 {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

//、
func moveZeroes_two(nums []int) {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[index] = nums[i]
			index++
		}
	}
	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}
}

// 等于0的都到右边,不考虑顺序
func moveZeroes_three(nums []int) {
	i, j := 0, len(nums)-1
	for i <= j {
		for i <= j && nums[i] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
		for i <= j && nums[i] != 0 {
			i++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes_two(nums)
	fmt.Println(nums)
}
