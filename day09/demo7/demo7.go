package main

import (
	"fmt"
)

func myAtoi(str string) int {
	if len(str) <= 0 {
		return 0
	}

	// 求出最小值和最大值
	min := -1 << 31
	max := 1<<31 - 1

	// 正负号的标记
	sign := 0

	// 初始值
	sum := 0
	strt := ""

	// 寻炸第一个非空字符
	i := 0
	for i < len(str) && str[i] == ' ' {
		i++
	}

	// 寻找数字
	rs := []rune(str)

	// 寻找负号
	for i < len(rs) {
		// 判断第一个字符是正数还是负数
		if i < len(rs) && (rs[i] == '-' || rs[i] == '+') {
			i++
			// 判断与正负号相连的下一个字符是数字还是其它
			if i < len(rs) && (rs[i]-'0' < 0 || rs[i]-'0' > 9) {
				return 0
			}
		}
		// 得出的必是数字，然后根据前一个判断符号，这里需要注意的是正好因为可以省略因此用了或的运算
		for i < len(rs) && rs[i]-'0' >= 0 && rs[i]-'0' <= 9 {
			if i == 0 || rs[i-1] == '+' {
				sign = 0
			} else if i > 0 && rs[i-1] == '-' {
				sign = -1
			}
			strt += string(rs[i])
			i++
		}

		break

	}

	for i := 0; i < len(strt); i++ {
		sum = sum*10 + int(strt[i]-'0')
		if sign == 0 && sum > max {
			return max
		} else if sign == -1 && (0-sum) <= min {
			return min
		}
	}
	if sign == -1 {
		return -sum
	}
	return sum
}
func main() {
	// fmt.Println('+')
	fmt.Println(myAtoi("+"))
}
