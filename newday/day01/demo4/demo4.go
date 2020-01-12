package main

import "fmt"

// 判断是否存在重复

func containsDuplicate(nums []int) bool {
	var m map[int]int
	m = make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return true
		} else {
			m[nums[i]] = i
		}
	}
	return false
}
func main() {
	data := []int{1, 2, 1, 4}
	fmt.Println(containsDuplicate(data))
}
