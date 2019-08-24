/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-17 13:42:58
 * @LastEditTime: 2019-08-17 15:36:01
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"fmt"
	"strings"
)

func Same(s string) bool {
	if strings.Contains(s[:len(s)-1], s[len(s)-1:]) {
		return true
	}
	return false
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	// 定义一个字符串数组存放结果
	result := s[:1]
	max := len(result)
	l := 0
	for i := 1; i < len(s); i++ {
		result += string(s[i])
		if Same(result) {
			l++
			i = l
			result = string(s[i])
		} else {
			if max < len(result) {
				max = len(result)
			}
		}
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring("asdgfdsg"))
}
