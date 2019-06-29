package solution

func combine(n int, k int) [][]int {
	var res [][]int
	var tmp []int
	backtrack(&res, tmp, 0, k, n)
	return res
}

func backtrack(res *[][]int, tmp []int, i int, k int, n int) {
	if len(tmp) == k {
		t := make([]int, k)
		copy(t, tmp)
		*res = append(*res, t)
	}
	for v := i + 1; v <= n && k-len(tmp) <= n-v+1; v++ {
		tmp = append(tmp, v)
		backtrack(res, tmp, v, k, n)
		tmp = tmp[:len(tmp)-1]
	}
}
