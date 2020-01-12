package main

import "fmt"

// 两个数组的交集 II
// 排序的情况
func intersect(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int, len(nums1))
	m2 := make(map[int]int, len(nums2))
	for i := 0; i < len(nums1); i++ {
		m1[nums1[i]]++
	}
	for j := 0; j < len(nums2); j++ {
		m2[nums2[j]]++
	}
	for i, _ := range m1 {
		if _, ok := m2[i]; ok {
			if m2[i] < m1[i] {
				m1[i] = m2[i]
			}
		} else {
			delete(m1, i)
		}
	}
	result := []int{}
	for i, v := range m1 {
		for v != 0 {
			result = append(result, i)
			v--
		}

	}
	return result
}

// 使用插入排序
func Insert_sort(data []int) {
	for i := 1; i < len(data); i++ {
		for j := i; j > 0; j-- {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
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

//使用排序进行优化
func intersect_two(nums1 []int, nums2 []int) []int {
	Insert_sort(nums1)
	Insert_sort(nums2)
	i, j := 0, 0

	result := []int{}

	for i < len(nums1) && j < len(nums2) {
		for i < len(nums1) && j < len(nums2) && nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			i++
			j++
		}
		if i < len(nums1) && j < len(nums2) && nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return result
}
func main() {
	nums1 := []int{2, 2, 3, 4, 0}
	nums2 := []int{1, 2, 1, 12, 4, 2}
	fmt.Println(intersect(nums1, nums2))
	fmt.Println(intersect_two(nums1, nums2))

}
