package solution

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) < 1 {
		return false
	}
	i, j, p, q := 0, 0, len(matrix)-1, len(matrix[0])-1
	return search(matrix, target, i, j, p, q)
}

func search(matrix [][]int, target int, i int, j int, p int, q int) bool {
	if i > p || j > q {
		return false
	}
	if i == p && j == q {
		return matrix[i][j] == target
	}
	mid := (j + q) / 2
	var k int
	for k = i; k <= p; k++ {
		if matrix[k][mid] == target {
			return true
		}
		if matrix[k][mid] > target {
			break
		}
	}
	ri, rj, rp, rq := i, mid+1, k-1, q
	li, lj, lp, lq := k, j, p, mid-1
	return search(matrix, target, ri, rj, rp, rq) || search(matrix, target, li, lj, lp, lq)
}
