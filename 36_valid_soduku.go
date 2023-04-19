package main

import "strconv"

func isValidSudoku(board [][]byte) bool {
	rowsVisited := [9][9]bool{}
	colsVisited := [9][9]bool{}
	gridVisited := [9][9]bool{}
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			cell := board[row][col]
			if cell == '.' {
				continue
			}
			//Cell element as number
			cellElement, _ := strconv.Atoi(string(cell))
			cellElement--
			//Calculate the grid index .
			//E.g
			//0,1,2,
			//3,4,5,
			//6,7,8
			gridIndex := col/3 + (row/3)*3
			if rowsVisited[row][cellElement] || colsVisited[col][cellElement] || gridVisited[gridIndex][cellElement] {
				return false
			}
			gridVisited[gridIndex][cellElement] = true
			rowsVisited[row][cellElement] = true
			colsVisited[col][cellElement] = true
		}
	}
	return true
}
