package solution

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) < 1 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	i, j := 0, m*n-1
	for i <= j {
		k := (i + j) / 2
		if matrix[k/n][k%n] == target {
			return true
		} else if matrix[k/n][k%n] > target {
			j = k - 1
		} else {
			i = k + 1
		}
	}
	return false
}
