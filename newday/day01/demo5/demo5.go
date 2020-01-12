package main

import "fmt"

// 出现一次的数字,切片实现
func singleNumber(nums []int) int {
	//定义一个二维数组
	var result map[int]int
	result = make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if _, ok := result[nums[i]]; ok {
			result[nums[i]]++
		} else {
			result[nums[i]] = 1
		}
	}
	for i, v := range result {
		if v == 1 {
			return i
		}
	}
	return 0
}

// 借助排序
func singleNumber_two(nums []int) int {
	Quick_Sort(nums, 0, len(nums)-1)

	if len(nums) <= 1 || nums[0] != nums[1] {
		return nums[0]
	}
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
			return nums[i]
		}
	}
	if nums[len(nums)-1] != nums[len(nums)-2] {
		return nums[len(nums)-1]
	} else {
		return 0
	}

}
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

// 根据异或运算的特点，相同的数字经过异或运算后结果为0，除单独出现一次的数字外，其他数字都是出现两次的，
// 那么这些数字经过异或运算后结果一定是0。而任何数字与0进行异或运算都是该数字本身。所以对数组所有元素进行异或运算，运算结果就是题目的答案。
func singleNumber_three(nums []int) int {
	num := 0
	for i := 0; i < len(nums); i++ {
		num = num ^ nums[i]
	}
	return num
}
func main() {
	data := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber_two(data))
}
