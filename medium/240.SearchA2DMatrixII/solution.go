package solution

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) < 1 || len(matrix[0]) < 1 {
		return false
	}
	// set the first position at top-right, if the target is greater than value then row++, if the target is less than value then col--
	row, col := 0, len(matrix[0])-1
	for col >= 0 && row < len(matrix) {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col--
		} else {
			row++
		}
	}
	return false
}
