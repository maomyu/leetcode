package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	// j:=0
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] && i+len(needle)-1 < len(haystack) && haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
func main() {
	fmt.Println(strStr("abb", "bb"))
}
