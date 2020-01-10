package main

import "fmt"

func longestPalindrome(s string) string {
	// 转换成数组
	r := []rune(s)
	i, j := 0, 0
	mid := 0
	length, max := 0, 0
	signi, signj := 0, 0
	for i >= 0 && j < len(r) && mid < len(r) {
		// 当子串为奇数的时候
		for i >= 0 && j < len(r) && r[i] == r[j] {
			length++
			i--
			j++
		}
		length = 2*length - 1
		// 判断大小
		if max < length {
			max = length
			signi = i + 1
			signj = j

		}
		length = 0
		// 当子串为偶数的时候
		i = mid
		j = mid
		for i >= 0 && j < len(r)-1 && r[i] == r[j+1] {
			length++
			i--
			j++
		}
		length = 2 * length
		if max < length {
			max = length
			signi = i + 1
			signj = j + 1
		}
		length = 0
		mid++
		i = mid
		j = mid
	}
	return s[signi:signj]
}
func main() {
	// longestPalindrome("cbbaaba")
	fmt.Println(longestPalindrome("cbbd"))
}
