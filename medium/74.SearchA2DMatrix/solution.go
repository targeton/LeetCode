package solution

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) < 1 || len(matrix[0]) < 1 {
		return false
	}
	i, j := 0, len(matrix)-1
	for i <= j {
		k := (i + j) / 2
		if matrix[k][0] == target {
			return true
		} else if matrix[k][0] > target {
			j = k - 1
		} else {
			i = k + 1
		}
	}
	if i == 0 {
		return false
	}
	p, q := 0, len(matrix[i-1])-1
	for p <= q {
		r := (p + q) / 2
		if matrix[i-1][r] == target {
			return true
		} else if matrix[i-1][r] > target {
			q = r - 1
		} else {
			p = r + 1
		}
	}
	return false
}
