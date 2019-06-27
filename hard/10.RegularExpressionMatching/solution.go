package solution

func isMatch(s string, p string) bool {
	cache := make([][]bool, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		cache[i] = make([]bool, len(p)+1)
	}
	// cache[len(s)][len(p)] means s == "", p == ""
	// cache[i][j] means s[i:],p[j:]
	for i := len(s); i >= 0; i-- {
		for j := len(p); j >= 0; j-- {
			if j == len(p) {
				cache[i][j] = i == len(s)
			} else {
				first := i < len(s) && (s[i] == p[j] || p[j] == '.')
				if j+1 < len(p) && p[j+1] == '*' {
					cache[i][j] = first && cache[i+1][j] || cache[i][j+2]
				} else {
					cache[i][j] = first && cache[i+1][j+1]
				}
			}
		}
	}
	return cache[0][0]
}
