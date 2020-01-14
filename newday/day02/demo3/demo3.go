package main

import "fmt"

// 找到它的第一个不重复的字符
func firstUniqChar_two(s string) int {
	m := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]] == 1 {
			return i
		}
	}
	return -1
}
func firstUniqChar(s string) int {
	rnum := [26]int{}
	for i := 0; i < len(s); i++ {
		rnum[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if rnum[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
func main() {
	fmt.Println(firstUniqChar_two("aqq"))
}
