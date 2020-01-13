package main

import "fmt"

// 两数之和双指针排序法
func twoSum(nums []int, target int) []int {
	Quick_Sort(nums, 0, len(nums)-1)

	i, j := 0, len(nums)-1
	sign := 0
	for i < j {
		if nums[i]+nums[j] > target {
			j--
		} else if nums[i]+nums[j] < target {
			i++
		} else {
			sign = 1
			break
		}
	}
	result := []int{}
	if sign == 1 {
		result = append(result, i, j)
	}
	return result
}

// 快速排序
func Quick_Sort(nums []int, from, to int) {
	mid := from
	i, j := from, to
	for i <= j {
		for mid <= j && nums[mid] <= nums[j] {
			j--
		}
		if mid <= j {
			nums[mid], nums[j] = nums[j], nums[mid]
			mid = j
		}
		for mid >= i && nums[mid] >= nums[i] {
			i++
		}
		if mid >= i {
			nums[mid], nums[i] = nums[i], nums[mid]
			mid = i
		}
	}
	if mid-from > 1 {
		Quick_Sort(nums, from, mid-1)
	}
	if to-mid > 1 {
		Quick_Sort(nums, mid+1, to)
	}
}

// 插入排序
func Insert_sort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}
func twoSum_first(nums []int, target int) []int {
	i, j := 0, 1
	result := []int{}
	for i < len(nums) {
		for j = i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i, j)
			}
		}
		i++
	}
	return result
}
func main() {
	nums := []int{3, 2, 4}
	fmt.Println(twoSum_first(nums, 6))
}
