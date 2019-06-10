package solution

func fib(n int) int {
	if n < 2 {
		return n
	}
	a := [][]int{{1, 1}, {1, 0}}
	r := pow(a, n-1)
	return r[0][0]
}

func climbStairs(n int) int {
	a := [][]int{{1, 1}, {1, 0}}
	r := pow(a, n)
	return r[0][0]
}

func pow(a [][]int, n int) [][]int {
	ret := [][]int{{1, 0}, {0, 1}}
	for {
		if n <= 0 {
			break
		}
		if n&1 == 1 {
			ret = multiply(ret, a)
		}
		n >>= 1
		a = multiply(a, a)
	}
	return ret
}

func multiply(a [][]int, b [][]int) [][]int {
	var result = make([][]int, 2)
	for i := 0; i < 2; i++ {
		result[i] = make([]int, 2)
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			result[i][j] = a[i][0]*b[0][j] + a[i][1]*b[1][j]
		}
	}
	return result
}
