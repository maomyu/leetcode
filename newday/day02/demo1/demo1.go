package main

// 反转字符串
func reverseString(s []byte) {
	lengh := len(s)
	for i := 0; i < lengh/2; i++ {
		s[i], s[lengh-1-i] = s[lengh-1-i], s[i]
	}
}
func main() {

}
