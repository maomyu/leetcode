package main

import "fmt"

// 有效的数独
func isValidSudoku(board [][]byte) bool {
	row := [9]map[byte]int{}
	col := [9]map[byte]int{}
	box := [9]map[byte]int{}

	for i := 0; i < 9; i++ {
		row[i] = make(map[byte]int, 9)
		col[i] = make(map[byte]int, 9)
		box[i] = make(map[byte]int, 9)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if _, ok := row[i][board[i][j]]; board[i][j] != '.' && ok {
				return false
			} else {
				row[i][board[i][j]] = 1
			}
			// 判断列
			if _, ok := col[j][board[i][j]]; board[i][j] != '.' && ok {
				return false
			} else {
				col[j][board[i][j]] = 1
			}
			box_index := (i/3)*3 + j/3
			// 判断盒子
			if _, ok := box[box_index][board[i][j]]; board[i][j] != '.' && ok {
				return false
			} else {
				box[box_index][board[i][j]] = 1
			}

		}
	}

	return true
}
func main() {
	data := [9][9]int{}
	fmt.Println(data[0])
	fmt.Println()
}
