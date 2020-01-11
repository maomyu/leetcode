package main

import "strconv"

import "fmt"

var m map[int]string

func intToRoman(num int) string {

	str := strconv.FormatInt(int64(num), 10)
	s := ""
	for i := 0; i < len(str); i++ {
		n := int(str[i] - '0')
		j := (len(str) - 1 - i)
		for j > 0 {
			n = n * 10
			j--
		}
		s += GetStr(n)
	}
	return s
}
func GetStr(num int) string {
	m = make(map[int]string, 24)
	m = map[int]string{
		1:    "I",
		2:    "II",
		3:    "III",
		4:    "IV",
		5:    "V",
		6:    "VI",
		7:    "VII",
		8:    "VIII",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		60:   "LX",
		70:   "LXX",
		80:   "LXXX",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		600:  "DC",
		700:  "DCC",
		800:  "DCCC",
		900:  "CM",
		1000: "M",
	}
	str := ""
	// 判断是否在map
	if _, ok := m[num]; ok {
		str += string(m[num])
	} else {
		s := strconv.FormatInt(int64(num), 10)
		if len(s) == 4 {
			n := num / 1000
			for n > 0 {
				str += "M"
				n--
			}
		} else if len(s) == 3 {
			n := num / 100
			for n > 0 {
				str += "C"
				n--
			}
		} else if len(s) == 2 {
			n := num / 10
			for n > 0 {
				str += "X"
				n--
			}
		}
	}
	return str
}

// 更好的思路
func intToRoman_two(num int) string {

	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	strs := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	index := 0
	str := ""
	for index < 13 {
		for num >= nums[index] {
			str += strs[index]
			num = num - nums[index]
		}
		index++
	}
	return str
}

// 罗马数字转整数
func romanToInt(s string) int {
	var ma map[string]int
	ma = make(map[string]int, 14)
	ma = map[string]int{
		"#":  0,
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10, "XL": 40, "L": 50, "XC": 90, "C": 100, "CD": 400, "D": 500, "CM": 900, "M": 1000,
	}
	s += "##"
	sum := 0
	// 定义两个指针
	i, j := 0, 1
	for j < len(s) {
		// 将ij进行合并
		if _, ok := ma[string(s[i])+string(s[j])]; ok {
			sum += ma[string(s[i])+string(s[j])]
			i += 2
			j += 2
		} else {
			sum += ma[string(s[i])]
			i++
			j++
		}
	}
	return sum
}

func main() {
	fmt.Println(intToRoman_two(1238))
	fmt.Println(romanToInt("I"))
}
