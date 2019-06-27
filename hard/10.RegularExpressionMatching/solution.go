package solution

var cache [][]int

func isMatch(s string, p string) bool {
	cache = make([][]int, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		cache[i] = make([]int, len(p)+1)
	}
	return dp(0, 0, s, p)
}

func dp(i int, j int, s string, p string) bool {
	if cache[i][j] != 0 {
		return cache[i][j] > 0
	}
	var ans bool
	if j == len(p) {
		ans = i == len(s)
	} else {
		first := i < len(s) && (s[i] == p[j] || p[j] == '.')
		if j+1 < len(p) && p[j+1] == '*' {
			ans = first && dp(i+1, j, s, p) || dp(i, j+2, s, p)
		} else {
			ans = first && dp(i+1, j+1, s, p)
		}
	}
	if ans {
		cache[i][j] = 1
	} else {
		cache[i][j] = -1
	}
	return ans
}
