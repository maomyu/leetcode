package main

import "fmt"

// 字符串转化成整数
func StrtoInt(str *string) int {
	// 转换层字符数组
	sr := []rune(*str)
	i := 0
	n := 0
	for i < len(sr) {
		n = n*10 + int(sr[i]-48)
		i++
	}
	return n
}
func multiply(num1 string, num2 string) string {
	// 定义一个矩阵
	// var result [][]int
	var r1, r2 []int
	sr1 := []rune(num1)
	sr2 := []rune(num2)
	i := 0
	for i < len(sr1) {
		r1 = append(r1, int(sr1[i]-48))
		i++
	}
	i = 0
	for i < len(sr2) {
		r2 = append(r2, int(sr2[i]-48))
		i++
	}
	// m := 0
	for i = len(r2) - 1; i >= 0; i-- {
		var r3 []int
		for j := 0; j < len(r1); j++ {
			// 计算sum
			r3 = append(r3, r2[i]*r1[j])
		}

		fmt.Println(r3)
	}
	// return strconv.Itoa(n1 * n2)f
	return ""

}
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
func LargeNumberMultiplication(a string, b string) (reuslt string) {
	if (a[0]-'0') == 0 || (b[0]-'0') == 0 {
		return "0"
	}
	a = Reverse(a)
	b = Reverse(b)

	var c []int
	c = make([]int, len(a)+len(b))

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			t := int((a[i] - '0') * (b[j] - '0'))
			c[i+j] += t
		}
	}
	// fmt.Println(c)
	// 表示进位
	var plus int = 0
	for i := 0; i < len(c); i++ {
		temp := c[i] + plus
		plus = 0
		if c[i] == 0 && i == len(c)-1 {
			if temp != 0 {
				reuslt += string(temp + '0')
			}
			break
		}
		// fmt.Println(c)
		if temp > 9 {
			// 得出进位值
			plus = temp / 10
			reuslt += string(temp - plus*10 + '0')
		} else {
			reuslt += string(temp + '0')
		}

	}

	return Reverse(reuslt)
}
func main() {
	fmt.Println(LargeNumberMultiplication("999", "0"))
}
