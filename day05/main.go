/*
 * @Author: your name
 * @Date: 2019-11-23 20:58:24
 * @LastEditTime: 2019-11-25 21:42:22
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \GoWebf:\leetcode\day05\main.go
 */
// 解题思路
// 首先字符串s1的排列的可能性应该是它的长度的阶乘，因为字符串长度可能为10000，所以找出所有排列情况是不太可能。
// 我们可以转换思路，不要关注排列的形式，而是关注排列中元素的数量关系，
// 比如aab，那么，转换为数量关系就是{a:2,b:1}，因为S1长度为3，所以我们的窗口长度也为3，
// 如果我们在S2的找到了这样一个窗口符合出现a的次数是两个，b是一个，那么S2就是包含S1的排列的。
package main

import (
	"fmt"
	"sort"
)

type StringScilce []rune

func (s StringScilce) Len() int {
	return len(s)
}
func (s StringScilce) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s StringScilce) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// 滑动窗口
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) == 0 || len(s2) == 0 {
		return false
	}
	sr1 := []rune(s1)
	sort.Sort(StringScilce(sr1))
	// 循环遍历长字符串
	for i, _ := range s2 {
		// 结束条件
		if i == len(s2)-len(s1) {
			break
		}
		// 开始滑动窗口,获得需要比较的字符串
		str := s2[i : i+len(s1)]
		// fmt.Println(str)
		sr2 := []rune(str)
		// 排序数组
		sort.Sort(StringScilce(sr2))
		if string(sr1) == string(sr2) {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(checkInclusion("", "eidbaooo"))
}
