package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	length := len(board)
	//二维数组line表示的是对应的行中是否有对应的数字，比如line[0][3]
	//表示的是第0行（实际上是第1行，因为数组的下标是从0开始的）是否有数字3

	var line [9][9]int
	var column [9][9]int
	var cell [9][9]int

	for i := 0; i < length; i++{
		for j := 0; j < length; j++ {
			if board[i][j] == '.' {
				continue
			}
			//num是当前格子的数字
			num := board[i][j] - '0' - 1
			//k是第几个单元格，9宫格数独横着和竖着都是3个单元格
			k := i / 3 * 3 + j / 3
			//如果当前数字对应的行和列以及单元格，只要一个有数字，说明冲突了，直接返回false。
			//举个例子，如果line[i][num]不等于0，说明第i（i从0开始）行有num这个数字。
			if line[i][num] != 0 || column[j][num] != 0 || cell[k][num] != 0 {
				return false
			}

			//表示第i行有num这个数字，第j列有num这个数字，对应的单元格内也有num这个数字
			line[i][num] =1 
			column[j][num] =1
			cell[k][num] = 1
		}
	}

	return true
}

func main() {
	jsonStr := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

	fmt.Println(isValidSudoku(jsonStr))
}
