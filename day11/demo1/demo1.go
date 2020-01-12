package main

import (
	"fmt"
)

/*
思路1:来一个字符准加到字符串中，
如：12,假设原来的字符串为 ""
遍历1对应的字符串,追加成["a","b","c"],
然后在["a","b","c"]基础上追加def,["ad","ae","af"。。。]
*/
// 定义一个列表map
var numstostring map[byte][]string = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

// 电话号码的组合
func letterCombinations(digits string) []string {
	str := []string{}
	if len(digits) <= 0 {
		return str
	}
	// 重要！！！！否则不会进入第二层循环
	str = append(str, "")
	for i := 0; i < len(digits); i++ {
		nums := numstostring[digits[i]]
		pre := str
		str = []string{}
		for _, v := range pre {
			for _, n := range nums {
				str = append(str, v+n)
			}
		}
	}
	return str
}

func main() {
	fmt.Println(letterCombinations("234"))
}
