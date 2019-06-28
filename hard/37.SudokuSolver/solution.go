package solution

func solveSudoku(board [][]byte) {
	var pos [][]int
	var res [][]byte
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				pos = append(pos, []int{i, j})
			}
		}
	}
	step, n := 0, len(pos)
	backtrack(&res, board, pos, step, n)
	copy(board, res)
}

func backtrack(res *[][]byte, board [][]byte, pos [][]int, step int, n int) {
	if step == n {
		for _, v := range board {
			var tmp []byte
			for _, b := range v {
				tmp = append(tmp, b)
			}
			*res = append(*res, tmp)
		}
		return
	}
	row, col := pos[step][0], pos[step][1]
	for v := byte('1'); v <= byte('9'); v++ {
		if isValid(board, row, col, v) {
			board[row][col] = v
			backtrack(res, board, pos, step+1, n)
			board[row][col] = '.'
		}
	}
}

func isValid(board [][]byte, row int, col int, value byte) bool {
	// check the row array has same value
	for i := 0; i < len(board[row]); i++ {
		if board[row][i] == value {
			return false
		}
	}
	// check the column array has same value
	for j := 0; j < len(board); j++ {
		if board[j][col] == value {
			return false
		}
	}
	// check 3*3 array has same value
	r, c := row/3, col/3
	for p := 0; p < 3; p++ {
		for q := 0; q < 3; q++ {
			if board[r*3+p][c*3+q] == value {
				return false
			}
		}
	}
	return true
}
