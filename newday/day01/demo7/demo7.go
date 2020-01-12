package main

import "fmt"

import "strconv"

//数组表示的数字加1，排除首位为0
func plusOne(digits []int) []int {
	i := 0
	for i = len(digits) - 1; i > 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			break
		}
	}

	if i == 0 && digits[0] == 9 {
		digits[0] = 0
		// 说明上述循环没有提前终止
		digits = append(digits, 1)
		for j := len(digits) - 1; j > 0; j-- {
			digits[j], digits[j-1] = digits[j-1], digits[j]
		}
	} else if i == 0 {
		digits[0]++
	}
	return digits
}

// 常规思路,不考虑溢出的情况,如果要考虑溢出，如果要考虑溢出，又将回到上述的思路的字符串模拟加法操作
func plusOne_1(digits []int) []int {
	sum := 0
	for i := 0; i < len(digits); i++ {
		sum = sum*10 + digits[i]
	}
	str := strconv.FormatInt(int64(sum+1), 10)
	digits = []int{}
	for _, v := range str {
		digits = append(digits, int(v-'0'))
	}
	return digits
}

// 不排除首尾为0的情况
func plusOne_2(digits []int) []int {
	if digits[0] != 0 {
		// 补个0
		digits = append(digits, 0)
		for j := len(digits) - 1; j > 0; j-- {
			digits[j], digits[j-1] = digits[j-1], digits[j]
		}
	}
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			break
		}
	}
	return digits
}
func main() {
	nums := []int{9, 9, 9}
	fmt.Println(plusOne_2(nums))
}
