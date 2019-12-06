/*
 * @Author: your name
 * @Date: 2019-11-27 09:44:04
 * @LastEditTime: 2019-11-27 10:01:21
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit

 * @FilePath: \GoWebf:\leetcode\day06\main.go
 */
//  #string到int
//  int,err:=strconv.Atoi(string)
//  #string到int64
//  int64, err := strconv.ParseInt(string, 10, 64)
//  #int到string
//  string:=strconv.Itoa(int)
//  #int64到string
//  string:=strconv.FormatInt(int64,10)

package main

import (
	"fmt"
	"math"
	"strconv"
)

// 整数翻转
func reverse(x int) int {
	sign := 0
	if x < 0 {
		x = -x
		sign = 1
	}
	// 将int先转换成string
	str := strconv.FormatInt(int64(x), 10)
	strr := []rune(str)
	var strrr []rune
	strrr = make([]rune, len(strr))
	index := 0
	for i := len(strr) - 1; i > -1; i-- {
		strrr[index] = strr[i]
		index += 1
	}
	a, err := strconv.Atoi(string(strrr))
	if err != nil {
		// fmt.Println()
		return 0
	}
	if a > math.MaxInt32 || a < math.MinInt32 {
		return 0
	}
	if sign == 1 {

		return -a
	}
	return a
}

// 在数字没有溢出的前提下，对于正数和负数，左移一位都相当于乘以2的1次方，左移n位就相当于乘以2的n次方。
func reverse_one(x int) (num int) {
	for x != 0 {
		fmt.Println(num)

		num = num*10 + x%10
		x = x / 10
	}
	MaxInt32 := 1<<31 - 1
	MinInt32 := -1 << 31
	if num > MaxInt32 || num < MinInt32 {
		return 0
	}
	return
}

func main() {
	var a int
	a = 123
	fmt.Println(reverse_one(a))
}
