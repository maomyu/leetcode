package main

import "fmt"

// 有效的字母异wei
func isAnagram(sa string, ta string) bool {
	s := []rune(sa)
	t := []rune(ta)
	if len(s) != len(t) {
		return false
	}
	Quick_Sort(s, 0, len(s)-1)
	Quick_Sort(t, 0, len(s)-1)
	if string(s) == string(t) {
		return true
	}
	return false

}
func isAnagram_w(sa string, ta string) bool {
	s := []rune(sa)
	t := []rune(ta)
	if len(s) != len(t) {
		return false
	}
	snum := make(map[rune]int, len(sa))
	for i := 0; i < len(s); i++ {
		snum[s[i]]++
	}
	fmt.Println(snum)
	for j := 0; j < len(t); j++ {
		if _, ok := snum[s[j]]; ok {
			snum[t[j]]--
		} else {
			return false
		}

	}
	fmt.Println(snum)
	for _, v := range snum {
		if v > 0 {
			return false
		}
	}
	return true

}

// 快速排序
func Quick_Sort(nums []rune, from, to int) {
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
func main() {
	fmt.Println(isAnagram_w("a", "b"))
}
