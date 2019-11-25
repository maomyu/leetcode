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
)

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) == 0 || len(s2) == 0 {
		return false
	}
	// 计算出窗口的长度
	l := len(s1)
	// 定义一个集合类型，存放字符和出现的次数
	var m map[rune]int
	m = make(map[rune]int, l)
	for _, v := range s1 {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}
	// fmt.Println(m)
	var test map[rune]int
	sign := 0
	test = make(map[rune]int, l)
	// 遍历s2，判断连续的三个数是否存在与m相同
	for i := 0; i < len(s2)-(l-1); i++ {
		for j := i; j <= i+l-1; j++ {

			if _, ok := test[rune(s2[j])]; ok {
				test[rune(s2[j])] += 1
			} else {
				test[rune(s2[j])] = 1
			}
		}
		// fmt.Println(s2[i])
		// 循环m
		for k, v := range m {
			if test[k] != v {
				sign = 1 //当前没有找到
				test = make(map[rune]int, l)
				sign = 0
				break
			}
			sign++
		}
		// fmt.Println(test)
		// fmt.Println(s2[i])
		if sign == len(m) {
			// fmt.Println(m)
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
}
