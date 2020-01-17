package main

import "fmt"

var dict = []string{"leet", "code"}

func wordBreak(s string, wordDict []string) bool {

	return word_Break(s, wordDict, 0)
}
func word_Break(s string, wordDict []string, start int) bool {
	if start == len(s) {
		return true
	}
	for end := start + 1; end <= len(s); end++ {
		fmt.Println(s[start:end])
		if Contains(wordDict, s[start:end]) && word_Break(s, wordDict, end) {
			return true
		}
	}
	return false
}
func Contains(dict []string, s string) bool {
	for _, v := range dict {
		if v == s {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(wordBreak("leetcode", dict))
}
