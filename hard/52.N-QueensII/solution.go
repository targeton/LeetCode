package solution

var positions [][]int

func totalNQueens(n int) int {
	positions = [][]int{}
	count := 0
	row, col := 0, 0
	for row < n {
		old, find := len(positions), false
		for c := col; c < n; c++ {
			if isNotUnderAttack(row, c) {
				positions = append(positions, []int{row, c})
				if row+1 == n {
					count++
					find = true
				} else {
					row++
					break
				}
			}
		}
		if len(positions) < 1 {
			break
		}
		if len(positions) == old || find {
			last := positions[len(positions)-1]
			row = last[0]
			col = last[1] + 1
			positions = positions[:len(positions)-1]
		} else {
			col = 0
		}
	}
	return count
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
