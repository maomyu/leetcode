package main

import "fmt"

// 困难解法
func convert(s string, numRows int) string {
	if len(s) <= 1 || numRows == 1 {
		return s
	}
	// 求出一个二维数组的列数
	col := ((numRows - 1) * len(s)) / (2*numRows - 2)
	col++
	var t2 [100][100]rune

	k := 0
	r := []rune(s)

	// 循环赋值i
	j := 0
	for j < col {
		if k >= len(r) {
			break
		}
		// 偶数列
		i := 0
		for j%(numRows-1) == 0 && i < numRows && k < len(r) {
			t2[i][j] = r[k]
			k++
			i++
		}
		i--
		j++
		for j%(numRows-1) != 0 && k < len(r) {
			i--
			t2[i][j] = r[k]
			j++
			k++
		}
	}
	str := ""
	for i := 0; i < numRows; i++ {
		for j := 0; j < col; j++ {
			if int(t2[i][j]) != 0 {
				str += string(t2[i][j])
			}
			// fmt.Print(string(t2[i][j]), " ")
		}
		// fmt.Println()
	}
	return str
}

// 根据输出的结果，没有必要非要Z字形
func convert_two(str string, l int) string {
	var m map[int][]rune
	m = make(map[int][]rune, l)
	r := []rune(str)
	k := 0
	for k < len(r) {
		for i := 0; k < len(r) && i < l; i++ {
			m[i] = append(m[i], r[k])
			k++
		}
		for j := l - 2; k < len(r) && j >= 1; j-- {
			m[j] = append(m[j], r[k])
			k++
		}
	}
	str = ""
	for i := 0; i < l; i++ {
		for j := 0; j < len(m[i]); j++ {
			str += string(m[i][j])
		}
	}
	return str
}
func main() {
	// fmt.Println(convert("abc", 2))
	fmt.Println(convert_two("", 1))
}
