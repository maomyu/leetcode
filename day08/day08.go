// 反转字符串里的单词
package main

import (
	"fmt"
)

// 思路1，倒叙遍历
func reverseStr(str string) string {
	if len(str) == 0 {
		return ""
	}
	// 定义首尾指针，将空格去掉
	head, tail := 0, len(str)-1
	for head < tail {
		// 终止条件
		if string(str[head]) != " " && string(str[tail]) != " " {
			break
		} else {
			if string(str[head]) == " " {
				head++
			}
			if string(str[tail]) == " " {
				tail--
			}
		}
	}
	if len(str) == 1 && string(str[0]) == " " {

		return ""
	}
	str = str[head : tail+1]

	head = 0
	ss := []string{}
	for i := 0; i < len(str); i++ {
		if string(str[i]) == " " {
			ss = append(ss, str[head:i])
			for string(str[i]) == " " {
				i++
			}
			head = i
		}
	}
	ss = append(ss, str[head:])
	str = ""
	for i := len(ss) - 1; i >= 0; i-- {
		str += ss[i] + " "
	}
	return str[:len(str)-1]
}
func main() {
	// fmt.Println(reverseStr("the sky is blue"))
	// fmt.Println(reverseStr("  hello world!  "))
	// fmt.Println(reverseStr("a good   example"))
	// fmt.Println(reverseStr("a good   e"))
	fmt.Println(reverseStr(" "))
	fmt.Println(reverseStr("a"))

}
