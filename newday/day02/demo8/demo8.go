package main

import (
	"fmt"
	"strconv"
)

/**
1.     1
2.     11
3.     21
4.     1211
5.     111221
*/
func countAndSay(n int) string {
	fpre := "1"
	for i := 2; i <= n; i++ {
		fnext := getstr(fpre)
		fpre = fnext
	}
	return fpre
}
func getstr(str string) string {
	result := ""
	count := 1
	for j := 1; j < len(str); j++ {
		if str[j] == str[j-1] {
			count++
		} else {
			result = result + strconv.FormatInt(int64(count), 10) + string(str[j-1])
			count = 1
		}
	}
	result = result + strconv.FormatInt(int64(count), 10) + string(str[len(str)-1])
	return result
}
func main() {
	// fmt.Println(getstr("1111"))
	fmt.Println(countAndSay(4))
}
