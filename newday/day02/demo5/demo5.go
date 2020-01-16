package main

import (
	"fmt"
)

// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
func isPalindrome(s string) bool {
	str := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			str += string(s[i])
		} else if s[i] >= 'a' && s[i] <= 'z' {
			str += string(s[i] - 32)
		} else if s[i] >= '0' && s[i] <= '9' {
			str += string(s[i])
		}
	}
	i, j := 0, len(str)-1
	for i <= j {
		if str[i] == str[j] {
			i++
			j--
		} else {
			return false
		}
	}
	// 将字符串进行转换
	return true
}
func main() {
	fmt.Println(isPalindrome("op"))
}
