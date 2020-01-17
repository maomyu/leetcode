package main

import "fmt"

import "strings"

func simplifyPath(strs string) string {
	result := strings.Split(strs, "/")
	str := []string{}
	for _, v := range result {
		if v == ".." {
			if len(str) > 0 {
				str = str[:len(str)-1]
			}
		} else if v != "." && v != " " {
			if len(v) != 0 {
				str = append(str, v)
			}
		}
	}
	re := "/"
	for i := 0; i < len(str); i++ {
		re = re + str[i] + "/"
	}
	if len(re) == 1 {
		return re
	}
	return re[:len(re)-1]
}
func main() {
	fmt.Println(simplifyPath("/home/"))
}
