package solution

import "strings"

func generateParenthesis(n int) []string {
	var res []string
	generate(&res, "", 0, n)
	return res
}

func generate(res *[]string, tmp string, i int, n int) {
	if i == 2*n {
		*res = append(*res, tmp)
		return
	}
	for _, v := range []string{"(", ")"} {
		if isValid(tmp, v, n) {
			tmp += v
			generate(res, tmp, i+1, n)
			tmp = tmp[:len(tmp)-1]
		}
	}
}

func isValid(tmp string, c string, n int) bool {
	if c == "(" && strings.Count(tmp, c) < n {
		return true
	}
	if c == ")" && strings.Count(tmp, ")") < strings.Count(tmp, "(") {
		return true
	}
	return false
}
