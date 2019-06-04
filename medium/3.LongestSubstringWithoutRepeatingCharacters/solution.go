package solution

func lengthOfLongestSubstring(s string) int {
	i, result, length := 0, 0, len(s)
	cache := make(map[byte]int)
	for j := 0; j < length; j++ {
		if pos, ok := cache[s[j]]; ok {
			if i < pos {
				i = pos
			}
		}

		if result < j+1-i {
			result = j + 1 - i
		}
		cache[s[j]] = j + 1
		if j >= length-1 {
			break
		}
	}
	return result
}
