package solution

import "strings"

func solveNQueens(n int) [][]string {
	var res [][]string
	var tmp []string
	backtrack(&res, tmp, 0, n)
	return res
}

func backtrack(res *[][]string, tmp []string, row int, n int) {
	if row == n {
		*res = append(*res, tmp)
		return
	}
	for col := 0; col < n; col++ {
		if canPlace(tmp, row, col, n) {
			s := strings.Repeat(".", n)
			ss := s[:col] + "Q" + s[col+1:]
			tmp = append(tmp, ss)
			backtrack(res, tmp, row+1, n)
			tmp = tmp[:len(tmp)-1]
		}
	}
}

func canPlace(tmp []string, row int, col int, n int) bool {
	if len(tmp) < 1 {
		return true
	}
	for i, value := range tmp {
		if value[col] == 'Q' {
			return false
		}
		if index := row + col - i; index >= 0 && index < n && value[index] == 'Q' {
			return false
		}
		if index := col + i - row; index >= 0 && index < n && value[index] == 'Q' {
			return false
		}
	}
	return true
}
