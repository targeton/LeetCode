package solution

import "strings"

func longestPalindrome(s string) string {
	ss := strings.Join(strings.Split(s, ""), "#")
	ss = "#" + ss + "#"
	p := make([]int, len(ss))
	id, mx := 0, 0
	for i := range ss {
		if i < mx {
			if p[2*id-i] < (mx - i) {
				p[i] = p[2*id-i]
			} else {
				p[i] = mx - i
			}
		} else {
			p[i] = 1
		}
		left, right := i-p[i], i+p[i]
		for {
			if left < 0 || right >= len(ss) {
				break
			}
			if ss[left] == ss[right] {
				p[i]++
				left--
				right++
			} else {
				break
			}
		}
		if i+p[i] > mx {
			id = i
			mx = i + p[i]
		}
	}
	sid, sMax := 0, 0
	for i := range p {
		if p[i] > sMax {
			sid = i
			sMax = p[i]
		}
	}
	sLen := sMax - 1
	if sLen%2 == 0 {
		return s[sid/2-sLen/2 : sid/2+sLen/2]
	}
	return s[sid/2-sLen/2 : sid/2+sLen/2+1]
}
