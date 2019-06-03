package solution

import "strings"

func lengthOfLongestSubstring(s string) int {
	i, j, result, length := 0, 0, 0, len(s)
	for i = 0; i < length; i++ {
		for j = i + 1; j < length; j++ {
			if strings.Contains(s[i:j], s[j:j+1]) {
				break
			}
		}
		if result < (j - i) {
			result = j - i
		}
	}
	return result
}
