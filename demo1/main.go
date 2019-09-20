/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-20 10:57:34
 * @LastEditTime: 2019-09-20 17:20:13
 * @LastEditors: Please set LastEditors
 */
package main

func dfs(M [][3]int, flag []bool, i int) {

	for index := 0; index < len(M); index++ {

		if flag[index] == false {
			flag[index] = true
			dfs(M, flag, index)

		}
	}
}

func main() {
	m := make([][3]int, 3)

	m[0][0] = 1
	m[0][1] = 1
	m[0][2] = 0
	m[1][0] = 1
	m[1][1] = 1
	m[1][2] = 0
	m[2][0] = 0
	m[2][1] = 0
	m[2][2] = 1
	flag := make([]bool, len(m))
	dfs(m, flag, 0)
}
