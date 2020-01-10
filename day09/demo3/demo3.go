package main

import "strconv"

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.FormatInt(int64(x), 10)
	i, j := 0, len(str)-1
	for i <= j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 将字符串进行反转后进行比较
func isPalindrome_two(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.FormatInt(int64(x), 10)
	s1 := ""
	for i := len(str) - 1; i >= 0; i-- {
		s1 += string(str[i])
	}
	if s1 == str {
		return true
	}
	return false
}
func main() {
	// fmt.Println(isPalindrome(123))
	// fmt.Println(isPalindrome(13))
	// fmt.Println(isPalindrome(121))
	// fmt.Println(isPalindrome(1221))
	// fmt.Println(isPalindrome(0))
	fmt.Println(isPalindrome_two(123))
	fmt.Println(isPalindrome_two(13))
	fmt.Println(isPalindrome_two(121))
	fmt.Println(isPalindrome_two(1221))
	fmt.Println(isPalindrome_two(0))
}
