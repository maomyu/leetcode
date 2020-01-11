package main

import (
	"fmt"
)

// 盛最多水的容器，双指针法
func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max := 0

	for i <= j {
		if height[i] > height[j] {
			temp := height[j] * (j - i)
			j--
			if temp >= max {
				max = temp
			}
		} else {
			temp := height[i] * (j - i)
			i++
			if temp >= max {
				max = temp
			}
		}
	}
	return max
}

func main() {
	data := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(data))
}
