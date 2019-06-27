package solution

func isMatch(s string, p string) bool {
	if len(p) < 1 {
		return len(s) < 1
	}
	first := len(s) > 0 && (s[0] == p[0] || p[0] == '.')
	if len(p) > 1 && p[1] == '*' {
		return first && isMatch(s[1:], p) || isMatch(s, p[2:])
	}
	return first && isMatch(s[1:], p[1:])
}
