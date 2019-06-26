package solution

import "strings"

var positions [][]int
var res [][]string

func solveNQueens(n int) [][]string {
	positions = [][]int{}
	res = [][]string{}
	backtrack(0, n)
	return res
}

func backtrack(row int, n int) {
	for col := 0; col < n; col++ {
		if isNotUnderAttack(row, col) {
			placeQueen(row, col)
			if row+1 == n {
				res = append(res, convert(positions))
			} else {
				backtrack(row+1, n)
			}
			removeQueen(row, col)
		}
	}
}

func isNotUnderAttack(row, col int) bool {
	if positions == nil || len(positions) < 1 {
		return true
	}
	for _, value := range positions {
		if row == value[0] || col == value[1] || (row-value[0]) == (col-value[1]) || (row+col) == (value[0]+value[1]) {
			return false
		}
	}
	return true
}

func placeQueen(row, col int) {
	positions = append(positions, []int{row, col})
}

func removeQueen(row, col int) {
	positions = positions[0 : len(positions)-1]
}

func convert(positions [][]int) []string {
	var result []string
	tmp := strings.Repeat(".", len(positions))
	for _, value := range positions {
		conv := tmp[:value[1]] + "Q" + tmp[value[1]+1:]
		result = append(result, conv)
	}
	return result
}
