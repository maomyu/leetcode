package main

import (
	"fmt"
)

// [4,0,5,-5,3,3,0,-4,-5]
// -2
func threeSumClosest(nums []int, target int) int {
	Quick_Sort(nums, 0, len(nums)-1)
	atemp := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums); i++ {
		start := i + 1
		end := len(nums) - 1
		for start < end {
			sum := nums[i] + nums[start] + nums[end]
			if getAbslute(target-sum) < getAbslute(target-atemp) {
				atemp = sum
			}
			if sum > target {
				end--
			} else if sum < target {
				start++
			} else {
				return atemp
			}

		}
	}
	return atemp
}

func getAbslute(nums int) int {
	if nums < 0 {

		return -nums
	}

	return nums
}

// 对数组进行快速排序
func Quick_Sort(nums []int, from int, to int) {
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
			nums[i], nums[mid] = nums[mid], nums[i]
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

// /[-4,-7,-2,2,5,-2,1,9,3,9,4,9,-9,-3,7,4,1,0,8,5,-7,-7]
// 29
func main() {
	data := []int{-4, -7, -2, 2, 5, -2, 1, 9, 3, 9, 4, 9, -9, -3, 7, 4, 1, 0, 8, 5, -7, -7}
	fmt.Println(threeSumClosest(data, 29))
}
