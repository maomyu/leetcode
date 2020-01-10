package main

import (
	"fmt"
)

func GetRightStr(str string) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	if len(str) <= 2 {
		return str
	}
	rstr := []rune(str)
	// 定义两个指针相差一位
	i, j := 0, 2
	for j < len(str) {
		// 判断指针指向的值是否相同
		for j < len(str) && rstr[i] == rstr[j] && rstr[j] == rstr[(i+j)/2] {
			rstr[(i+j)/2] = '#'
			rstr[i], rstr[(i+j)/2] = rstr[(i+j)/2], rstr[i]
			i++
			j++
		}
		for j < len(str)-1 && rstr[i] != rstr[j] && rstr[i+1] == rstr[i] && rstr[j+1] == rstr[j] {
			rstr[j] = '#'
			rstr[i], rstr[j] = rstr[j], rstr[i]
			i++
			j++
		}
		i++
		j++
	}
	str = ""
	for i := 0; i < len(rstr); i++ {
		if rstr[i] != '#' {
			str += string(rstr[i])
		}
	}
	return str
}

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	str1 := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s\n", &str1[i])
	}
	for i := 0; i < n; i++ {
		fmt.Println(GetRightStr(str1[i]))
	}

}
