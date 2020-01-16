package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	// 假设数组中长度最小的字符串为公共长度
	minstr := strs[0]

	for _, v := range strs {
		if len(minstr) >= len(v) {
			minstr = v
		}
	}
	// fmt.Println(minstr)
	for _, v := range strs {
		minstr = findminstr(v, minstr)
		if len(minstr) == 0 {
			return ""
		}
	}
	return minstr
}
func findminstr(str string, min string) string {
	str = str[:len(min)]
	for i, v := range str {
		if v != rune(min[i]) {
			min = str[:i]
			break
		}
	}
	return min
}

func main() {

}
