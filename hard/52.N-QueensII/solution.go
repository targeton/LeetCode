package solution

var positions [][]int

func totalNQueens(n int) int {
	positions = [][]int{}
	count := 0
	backtrack(0, n, &count)
	return count
}

func backtrack(row int, n int, count *int) {
	for col := 0; col < n; col++ {
		if isNotUnderAttack(row, col) {
			placeQueen(row, col)
			if row+1 == n {
				*count++
			} else {
				backtrack(row+1, n, count)
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
